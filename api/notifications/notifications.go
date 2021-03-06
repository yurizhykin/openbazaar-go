package notifications

import (
	"encoding/json"
	"fmt"
	"time"
)

type Notification struct {
	ID        int       `json:"id"`
	Data      Data      `json:"notification"`
	Timestamp time.Time `json:"timestamp"`
	Read      bool      `json:"read"`
}

type Data interface {
	// TODO maybe should be made 'real interface', which will allow
	// to use typed channels, type checking and semantic dispatching
	// instead of typecase:

	// Serialize() []byte
	// Describe() (string, string)
}

type notificationWrapper struct {
	Notification Data `json:"notification"`
}

type messageWrapper struct {
	Message Data `json:"message"`
}

type messageReadWrapper struct {
	MessageRead interface{} `json:"messageRead"`
}

type messageTypingWrapper struct {
	MessageRead interface{} `json:"messageTyping"`
}

type orderWrapper struct {
	OrderNotification `json:"order"`
}

type paymentWrapper struct {
	PaymentNotification `json:"payment"`
}

type orderConfirmationWrapper struct {
	OrderConfirmationNotification `json:"orderConfirmation"`
}

type orderCancelWrapper struct {
	OrderCancelNotification `json:"orderConfirmation"`
}

type refundWrapper struct {
	RefundNotification `json:"refund"`
}

type fulfillmentWrapper struct {
	FulfillmentNotification `json:"orderFulfillment"`
}

type completionWrapper struct {
	CompletionNotification `json:"orderCompletion"`
}

type disputeOpenWrapper struct {
	DisputeOpenNotification `json:"disputeOpen"`
}

type disputeUpdateWrapper struct {
	DisputeUpdateNotification `json:"disputeUpdate"`
}

type disputeCloseWrapper struct {
	DisputeCloseNotification `json:"disputeClose"`
}

type OrderNotification struct {
	Title             string `json:"title"`
	BuyerGuid         string `json:"buyerGuid"`
	BuyerBlockchainId string `json:"buyerBlockchainId"`
	Thumbnail         string `json:"thumbnail"`
	Timestamp         int    `json:"timestamp"`
	OrderId           string `json:"orderId"`
}

type PaymentNotification struct {
	OrderId      string `json:"orderId"`
	FundingTotal uint64 `json:"fundingTotal"`
}

type OrderConfirmationNotification struct {
	OrderId string `json:"orderId"`
}

type OrderCancelNotification struct {
	OrderId string `json:"orderId"`
}

type RefundNotification struct {
	OrderId string `json:"orderId"`
}

type FulfillmentNotification struct {
	OrderId string `json:"orderId"`
}

type CompletionNotification struct {
	OrderId string `json:"orderId"`
}

type DisputeOpenNotification struct {
	OrderId string `json:"orderId"`
}

type DisputeUpdateNotification struct {
	OrderId string `json:"orderId"`
}

type DisputeCloseNotification struct {
	OrderId string `json:"orderId"`
}

type FollowNotification struct {
	Follow string `json:"follow"`
}

type UnfollowNotification struct {
	Unfollow string `json:"unfollow"`
}

type ModeratorAddNotification struct {
	ModeratorAdd string `json:"moderatorAdd"`
}

type ModeratorRemoveNotification struct {
	ModeratorRemove string `json:"moderatorRemove"`
}

type StatusNotification struct {
	Status string `json:"status"`
}

type ChatMessage struct {
	MessageId string    `json:"messageId"`
	PeerId    string    `json:"peerId"`
	Subject   string    `json:"subject"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

type ChatRead struct {
	MessageId string `json:"messageId"`
	PeerId    string `json:"peerId"`
	Subject   string `json:"subject"`
}

type ChatTyping struct {
	PeerId  string `json:"peerId"`
	Subject string `json:"subject"`
}

func Serialize(i interface{}) []byte {
	var n notificationWrapper
	switch i.(type) {
	case OrderNotification:
		n = notificationWrapper{
			orderWrapper{
				OrderNotification: i.(OrderNotification),
			},
		}
	case PaymentNotification:
		n = notificationWrapper{
			paymentWrapper{
				PaymentNotification: i.(PaymentNotification),
			},
		}
	case OrderConfirmationNotification:
		n = notificationWrapper{
			orderConfirmationWrapper{
				OrderConfirmationNotification: i.(OrderConfirmationNotification),
			},
		}
	case OrderCancelNotification:
		n = notificationWrapper{
			orderCancelWrapper{
				OrderCancelNotification: i.(OrderCancelNotification),
			},
		}
	case RefundNotification:
		n = notificationWrapper{
			refundWrapper{
				RefundNotification: i.(RefundNotification),
			},
		}
	case FulfillmentNotification:
		n = notificationWrapper{
			fulfillmentWrapper{
				FulfillmentNotification: i.(FulfillmentNotification),
			},
		}
	case CompletionNotification:
		n = notificationWrapper{
			completionWrapper{
				CompletionNotification: i.(CompletionNotification),
			},
		}
	case DisputeOpenNotification:
		n = notificationWrapper{
			disputeOpenWrapper{
				DisputeOpenNotification: i.(DisputeOpenNotification),
			},
		}
	case DisputeUpdateNotification:
		n = notificationWrapper{
			disputeUpdateWrapper{
				DisputeUpdateNotification: i.(DisputeUpdateNotification),
			},
		}
	case DisputeCloseNotification:
		n = notificationWrapper{
			disputeCloseWrapper{
				DisputeCloseNotification: i.(DisputeCloseNotification),
			},
		}
	case FollowNotification:
		n = notificationWrapper{
			i.(FollowNotification),
		}
	case UnfollowNotification:
		n = notificationWrapper{
			i.(UnfollowNotification),
		}
	case ModeratorAddNotification:
		n = notificationWrapper{
			i.(ModeratorAddNotification),
		}
	case ModeratorRemoveNotification:
		n = notificationWrapper{
			i.(ModeratorRemoveNotification),
		}
	case StatusNotification:
		s := i.(StatusNotification)
		b, _ := json.Marshal(s)
		return b
	case ChatMessage:
		m := messageWrapper{
			i.(ChatMessage),
		}
		b, _ := json.MarshalIndent(m, "", "    ")
		return b
	case ChatRead:
		m := messageReadWrapper{
			i.(ChatRead),
		}
		b, _ := json.MarshalIndent(m, "", "    ")
		return b
	case ChatTyping:
		m := messageTypingWrapper{
			i.(ChatTyping),
		}
		b, _ := json.MarshalIndent(m, "", "    ")
		return b
	case []byte:
		return i.([]byte)
	}

	b, _ := json.MarshalIndent(n, "", "    ")
	return b
}

func Describe(i interface{}) (string, string) {
	var head, body string
	switch i.(type) {
	case OrderNotification:
		head = "Order received"

		n := i.(OrderNotification)
		var buyer string
		if n.BuyerBlockchainId != "" {
			buyer = n.BuyerBlockchainId
		} else {
			buyer = n.BuyerGuid
		}
		form := "You received an order \"%s\".\n\nOrder ID: %s\nBuyer: %s\nThumbnail: %s\nTimestamp: %d"
		body = fmt.Sprintf(form, n.Title, n.OrderId, buyer, n.Thumbnail, n.Timestamp)

	case PaymentNotification:
		head = "Payment received"

		n := i.(PaymentNotification)
		form := "Payment for order \"%s\" received (total %d)."
		body = fmt.Sprintf(form, n.OrderId, n.FundingTotal)

	case OrderConfirmationNotification:
		head = "Order confirmed"

		n := i.(OrderConfirmationNotification)
		form := "Order \"%s\" has been confirmed."
		body = fmt.Sprintf(form, n.OrderId)

	case OrderCancelNotification:
		head = "Order cancelled"

		n := i.(OrderCancelNotification)
		form := "Order \"%s\" has been cancelled."
		body = fmt.Sprintf(form, n.OrderId)

	case RefundNotification:
		head = "Payment refunded"

		n := i.(RefundNotification)
		form := "Payment refund for order \"%s\" received."
		body = fmt.Sprintf(form, n.OrderId)

	case FulfillmentNotification:
		head = "Order fulfilled"

		n := i.(FulfillmentNotification)
		form := "Order \"%s\" was marked as fulfilled."
		body = fmt.Sprintf(form, n.OrderId)

	case CompletionNotification:
		head = "Order completed"

		n := i.(CompletionNotification)
		form := "Order \"%s\" was marked as completed."
		body = fmt.Sprintf(form, n.OrderId)

	case DisputeOpenNotification:
		head = "Dispute opened"

		n := i.(DisputeOpenNotification)
		form := "Dispute around order \"%s\" was opened."
		body = fmt.Sprintf(form, n.OrderId)

	case DisputeUpdateNotification:
		head = "Dispute updated"

		n := i.(DisputeUpdateNotification)
		form := "Dispute around order \"%s\" was updated."
		body = fmt.Sprintf(form, n.OrderId)

	case DisputeCloseNotification:
		head = "Dispute closed"

		n := i.(DisputeCloseNotification)
		form := "Dispute around order \"%s\" was closed."
		body = fmt.Sprintf(form, n.OrderId)
	}
	return head, body
}
