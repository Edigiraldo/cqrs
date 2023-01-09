# CQRS

To run the application:

```console
  docker-compose up -d --build
```

Open a WebSocket connection, every time a resource is created you will be notified:

```console
  localhost:8080/ws
```

Create some feeds:

- [POST] localhost:8080/feeds

  ```json
  {
    "title": "The Fascinating World of Quantum Computing: How It Works and Its Potential Impacts",
    "description": "Learn about the principles of quantum computing and how it could revolutionize various industries"
  }
  ```

- [POST] localhost:8080/feeds

  ```json
  {
    "title": "The Top 10 Healthy Habits for a Long and Happy Life",
    "description": "Discover simple habits to incorporate into your daily routine for improved health and well-being"
  }
  ```

- [POST] localhost:8080/feeds

  ```json
  {
    "title": "Mastering Git for Version Control and Collaboration",
    "description": "Learn how to use Git for version control and collaboration in software development"
  }
  ```

- [POST] localhost:8080/feeds

  ```json
  {
    "title": "The Best Places to Visit in South America: A Guide to the Continent's Top Destinations",
    "description": "Find a list of the top places to visit in South America and a
    ```dd them to your travel bucket list"
  }

- [POST] localhost:8080/feeds

  ```json
  {
    "title": "The Future of Work: How Artificial Intelligence and Automation Are Changing the Landscape",
    "description": "Learn about the ways in which artificial intelligence and a
    ```utomation are changing the world of work and what it means for the future"
  }

- [POST] localhost:8080/feeds

  ```json
  {
    "title": "The Benefits of Plant-Based Eating for Health and the Environment",
    "description": "Learn about the many benefits of a plant-based diet and how to incorporate more plant-based meals into your diet"
  }
  ```

- [POST] localhost:8080/feeds

  ```json
  {
    "title": "The Top 10 Hiking Trails in the United States",
    "description": "Discover the best hiking trails in the United States and add them to your outdoor adventure list"
  }
  ```

- [POST] localhost:8080/feeds

  ```json
  {
    "title": "The Science of Sleep: How It Affects Your Health and Well-Being",
    "description": "Learn about the importance of sleep and how to improve your sleep habits for better health and well-being"
  }
  ```

- [POST] localhost:8080/feeds

  ```json
  {
    "title": "The Economic Impact of the COVID-19 Pandemic: A Global Perspective",
    "description": "Understand the ways in which the COVID-19 pandemic has impacted the global economy and what it means for the future"
  }
  ```

- [POST] localhost:8080/feeds

  ```json
  {
    "title": "The Top 5 Apps for Improved Mental Health and Wellness",
    "description": "Find a list of the top mental health and wellness apps to help improve your mental well-being"
  }
  ```

Get all the feed you have created:

- [GET] localhost:8080/feeds

And now search feeds that have any topic of your interest:

- localhost:8080/search?q=softwar
