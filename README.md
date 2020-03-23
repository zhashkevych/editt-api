# editt API 

## Requirements
- go 1.13
- docker & docker-compose

## Run Project

Use ```make run-local``` to build and run docker containers with application itself and mongodb instance

Use ```make run-deploy``` on production environment to run API container

## Client's API:

### GET /api/publications

Retrieves all publications hosted on platform:

#### Example Params: 
```
/api/publications?type=popular&limit=3
```

Where: type (popular/latest) and limit is arbitrary int

#### Example Response: 
```
{
   "publications":[
      {
         "id":"5e7202199aa5d6d4014d4993",
         "author":"Maksim Zhashkevych",
         "title":"Стань Богаче, Используя этот Простой Трюк",
         "tags":[
            "финансы",
            "деньги",
            "бюджет"
         ],
         "body":"text",
         "imageLink":"https://miro.medium.com/max/10680/0*Jz9bUxFLDCZEGlxb",
         "reactions":0,
         "views":1,
         "readingTime":2,
         "publishedAt":"2020-03-18T11:12:25.583Z"
      }
   ]
}
```

### POST /api/publications

Used to publish new publication

##### Example Input: 
```
{
	"author": "Maksim Zhashkevych",
	"title": "Стань Богаче, Используя этот Простой Трюк", 
	"tags": ["финансы", "деньги", "бюджет"], 
	"body": "<p>text</p>",
	"imageLink": "https://miro.medium.com/max/10680/0*Jz9bUxFLDCZEGlxb"
	
}
```

### GET /api/publications/:id

Used to retrieve single publication

##### Example Response: 
```
{
    "id": "5e7202199aa5d6d4014d4993",
    "author": "Maksim Zhashkevych",
    "title": "Стань Богаче, Используя этот Простой Трюк",
    "tags": [
        "финансы",
        "деньги",
        "бюджет"
    ],
    "body": "<p>text</p>",
    "imageLink": "https://miro.medium.com/max/10680/0*Jz9bUxFLDCZEGlxb",
    "reactions": 0,
    "views": 2,
    "readingTime": 2,
    "publishedAt": "2020-03-18T11:12:25.583Z"
}
```

### POST /api/publications/:id/reaction

Used to increse reactions count for specific publication

## Admin's Panel API:

### POST /admin/sign-in

Used to get authorization token

##### Example Input: 
```
{
	"username": "edittor",
	"password": "edittor"
} 
```

##### Example Response: 
```
{
	"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzEwMzgyMjQuNzQ0MzI0MiwidXNlciI6eyJJRCI6IjAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMCIsIlVzZXJuYW1lIjoiemhhc2hrZXZ5Y2giLCJQYXNzd29yZCI6IjQyODYwMTc5ZmFiMTQ2YzZiZDAyNjlkMDViZTM0ZWNmYmY5Zjk3YjUifX0.3dsyKJQ-HZJxdvBMui0Mzgw6yb6If9aB8imGhxMOjsk"
} 
```

### GET /admin/metrics

Used to retrieve metrics

##### Example Response: 
```
{
    "last24": [
        {
            "unique_visitors_count": 245,
            "timestamp": "2020-03-23T15:25:01.908Z"
        },
        {
            "unique_visitors_count": 258,
            "timestamp": "2020-03-23T15:22:20.514Z"
        }
    ],
    "publications_count": 3
}
```

### GET /admin/publications

Retrieves all platform publications

##### Example Response: 
```
{
   "publications":[
      {
         "id":"5e7202199aa5d6d4014d4993",
         "author":"Maksim Zhashkevych",
         "title":"Стань Богаче, Используя этот Простой Трюк",
         "tags":[
            "финансы",
            "деньги",
            "бюджет"
         ],
         "body":"text",
         "imageLink":"https://miro.medium.com/max/10680/0*Jz9bUxFLDCZEGlxb",
         "reactions":0,
         "views":1,
         "readingTime":2,
         "publishedAt":"2020-03-18T11:12:25.583Z"
      }
   ]
}
```

### DELETE /admin/publications/:id

Used to remove publication by ID