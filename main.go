package main

import (
	"fmt"
	"strings"
	"time"
)

var knowledgeBase = map[string]string{
	"order_status":    "You can check your order status on our website under 'My Orders'.",
	"track_order":     "Please go to the 'Track Order' section and enter your tracking number.",
	"return_policy":   "Our return policy allows for returns within 30 days of purchase.",
	"contact_support": "You can contact our support team via email at support@example.com or call us at 1-800-SUPPORT.",
}

var requestQueue = make(chan string, 10)

func identifyIntent(query string) string {
	query = strings.ToLower(query)
	if strings.Contains(query, "order status") || strings.Contains(query, "where is my order") {
		return "order_status"
	} else if strings.Contains(query, "track order") || strings.Contains(query, "tracking number") {
		return "track_order"
	} else if strings.Contains(query, "return policy") || strings.Contains(query, "return an item") {
		return "return_policy"
	} else if strings.Contains(query, "contact support") || strings.Contains(query, "reach customer service") {
		return "contact_support"
	}
	return "unknown"
}

func AIAssistant() {
	fmt.Println("AI Agent started. Listening for queries...")
	for query := range requestQueue {
		fmt.Printf("AI Agent received query: %s\n", query)
		intent := identifyIntent(query)
		fmt.Printf("Identified intent: %s\n", intent)

		if answer, found := knowledgeBase[intent]; found {
			fmt.Printf("AI Agent found answer: %s\n", answer)
			fmt.Printf("Response sent: %s\n\n", answer)
		} else {
			response := "Sorry, I don't have information on that. Please try rephrasing."
			fmt.Printf("AI Agent could not find answer.\n")
			fmt.Printf("Response sent: %s\n\n", response)
		}
		time.Sleep(time.Second) // simulate processing time
	}
	fmt.Println("AI Agent stopped.")
}

func simulateCustomerQuery(query string) {
	fmt.Printf("Customer asked: %s\n", query)
	requestQueue <- query
}

func main() {
	go AIAssistant()

	// test queries
	simulateCustomerQuery("What is my order status?")
	simulateCustomerQuery("How do I track my order?")
	simulateCustomerQuery("Tell me about your return policy.")
	simulateCustomerQuery("What is the weather like?")
	simulateCustomerQuery("I need to contact support.")

	// keep main function running for a while to allow the agent to process queries (in a real system, this wouldn't happen)
	time.Sleep(5 * time.Second)

	// close the request queue to signal the agent to stop (in a real system, this wouldn't happen)
	close(requestQueue)
	time.Sleep(time.Second)
	fmt.Println("Simulation finished.")
}
