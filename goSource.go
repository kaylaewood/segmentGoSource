package main

import "gopkg.in/segmentio/analytics-go.v3"

func main() {
  client := analytics.New("EOywUtBfG2gyQoTPKkMzUHrg8TzA7Dsv")
  defer client.Close()


  client.Enqueue(analytics.Identify{
    UserId: "019mr8mf4r",
    Traits: analytics.NewTraits().
      SetName("Kayla Wood").
      SetEmail("kwood@example.com").
      Set("plan", "Enterprise").
      Set("location", "Denver"),
  })
}
