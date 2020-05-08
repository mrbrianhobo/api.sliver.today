# 🍕 api.sliver.today 

## ℹ️ About  [![Build Status](https://travis-ci.com/mrbrianhobo/api.sliver.today.svg?branch=master)](https://travis-ci.com/mrbrianhobo/api.sliver.today)  [![LICENSE](https://img.shields.io/github/license/mrbrianhobo/api.sliver.today)](https://github.com/mrbrianhobo/api.sliver.today/blob/master/LICENSE)

[api.sliver.today](http://api.sliver.today) is a free REST API for retrieving [Sliver Pizzeria](https://www.sliverpizzeria.com)'s pizza of the day.

## 👀 Demo

➡️ **[sliver.today](http://sliver.today)**

## 🚀 Getting Started

### Try it out!

```bash
# Retrieve today's pizza menu
$ curl --request GET \
       --url 'https://api.sliver.today' \
       --header 'accept: application/json'

# Retrieve today's pizza menu by location
$ curl --request GET \
       --url 'https://api.sliver.today/?location={location}' \
       --header 'accept: application/json'
```

### Running Locally

```bash
$ make install
$ make dev
go run main.go
2020/05/07 18:17:00 Listening on localhost:8080
```

Now you can visit [http://localhost:8080](http://localhost:8080) and test out the API!  
TODO: It's also bundled as a Docker image (for Google Cloud Run) so you can also run it that way. 

### Pizza Menu Object

```go
type Menu struct {
  Location string `json:"location"`
  Date     string `json:"date"`
  Pizza    string `json:"pizza"`
}
```

`Location: string`: The location of the specified menu (e.g. `"telegraph"`, `"shattuck"`, or `"broadway"`).

`Date: string`: The date of the requested menu. In the format `YYYY-MM-DD`.

`Pizza: string`:  The pizza of the day for a specified location. (e.g. `"Corn, Roasted Yukon Gold Potatoes, Baby Spinach Mixed with Caramelized Onions, Mozzarella, French Feta Cheese, with Avocado Spread."`).

## ⚙️ Usage

### Request

Retrieves a list of pizza menu objects.

```jsx
GET https://api.sliver.today
```

### Response

```json
{
  "pizzas": [
    {
      "location": "telegraph",
      "date": "2020-05-06",
      "pizza": "Corn, Roasted Yukon Gold Potatoes, Baby Spinach Mixed with Caramelized Onions, Mozzarella, French Feta Cheese, with Avocado Spread."
    },
    {
      "location": "shattuck",
      "date": "2020-05-06",
      "pizza": "Wild Mushrooms (Shitake, Chanterelle, Portabella, Cremini Mushrooms), Mozzarella, Asiago Fresh, Green Scallons, Fresh Herbs, Infused Chanterelle Mushrooms, and Garlic Olive Oil."
    },
    {
      "location": "broadway",
      "date": "2020-05-06",
      "pizza": "Roasted Yukon Gold Potatoes, Baby Heirloom Tomatoes, Baby Spinach Mixed with Caramelized Onions, Mozzarella, Pecorino Cheese mixed with Fresh Herbs, and Infused Sage Garlic Oil."
    }
  ]
}
```

### Query Parameters

### `location?: string`

The location of the wanted pizza menu. Defaults to `all`. Supports `telegraph`, `shattuck`, `broadway` as parameters.

## 🤔 FAQ

**Can I use this in my project?**

Yup! Do whatever you want with it! Let me know if you're thinking of making something cool with it! 😎

**Is it open source? Can I contribute?**

Yes! Checkout my spaghetti code on Github and submit issues or bugs. 🤖[API](https://github.com/mrbrianhobo/api.sliver.today) 👨‍💻[Site](https://github.com/mrbrianhobo/sliver.today)

**What else do you want to do with this?** 

I'm thinking about making this into a Slackbot/Discord Bot/Twitter Bot. It would also be pretty cool to index all the past pizzas that Sliver has served. 

**What did you use to make this?**

Backend API: Golang + [Colly](http://go-colly.org) (🚀deployed on [Google Cloud Run](https://cloud.google.com/run))  
Frontend: Typescript + [Next.js](https://nextjs.org) + [MDX](https://mdxjs.com) (🎨styled with [Theme-UI](https://theme-ui.com) + [Emotion](https://emotion.sh))

**Why did you make this?**

IDK. I really like Sliver. And I was bored. Really bored. 😴

## 📝 License

MIT Licensed.
