# The Feedback Game

## API

### POST /api/login

Log in using your LDAP credentials.

Request example:

```
POST /api/login
{
    "username": "arturo.martinez",
    "password": "ILoveBigBananas"
}
```

Response example:

```
{
	"id": 0,
	"username": "arturo.martinez",
	"firstName": "Arturo",
	"surname": "Martinez Minguez",
	"fullName": "Arturo Martinez Minguez",
	"jobTitle": "Node.js Developer",
	"department": "Development",
	"company": "ICEMOBILE",
	"email": "arturo.martinez@icemobile.com",
	"avatar": "arturo.martinez.png",
	"role": "admin"
}
```

In the response you will find the header `x-auth-token`, which you have to send in every subsequent request. This header does not change its value.

* Role required: none

### GET /api/my/reviews

List all the reviews a user has assigned.

Response example:

```
[{
	"id": 31,
	"uuid": "91a637eb-083c-460c-92ca-9ab55c528b44",
	"reviewer": {
		"id": 1,
		"username": "arturo.martinez",
		"firstName": "Arturo",
		"surname": "Martinez Minguez",
		"fullName": "Arturo Martinez Minguez",
		"jobTitle": "Node.js Developer",
		"department": "Development",
		"company": "ICEMOBILE",
		"email": "arturo.martinez@icemobile.com",
		"avatar": "arturo.martinez.png",
		"role": ""
	},
	"reviewee": {
		"id": 4,
		"username": "andrew.gerssen",
		"firstName": "Andrew",
		"surname": "Gerssen",
		"fullName": "Andrew Gerssen",
		"jobTitle": "Delivery Manager",
		"department": "Project Management",
		"company": "ICEMOBILE",
		"email": "andrew.gerssen@icemobile.com",
		"avatar": "andrew.gerssen.png",
		"role": ""
	},
	"cards": null,
	"remark": "Lorem ipsum",
	"completed": false
}]
```

* Role required: user

### GET /api/reviews

Get all the reviews. Only used in the admin console.

* Role required: user

### POST /api/reviews

Create a review. Only used in the admin console.

* Role required: admin

### GET /api/reviews/:id

Get a single review. Only used in the admin console.

* Role required: user

### PUT /api/reviews/:id

Update a review. Under construction.

* Role required: user

### DELETE /api/reviews/:id

Delete a review. Only used in the admin console.

* Role required: admin

### GET /api/cards

Get all the cards.

Example response:

```
[{
	"id": 20,
	"title": "Accurate",
	"category": 0 // 0: Positive
}, {
	"id": 39,
	"title": "Always talking",
	"category": 1 // 1: Negative
}, {
	"id": 16,
	"title": "Ambitious",
	"category": 0
}]
```

* Role required: user

### POST /api/cards

Create a card. Only used in the admin console.

* Role required: admin

### PUT /api/cards/:id

Update a card. Only used in the admin console.

* Role required: admin

### DELETE /api/cards/:id

Delete a card. Only used in the admin console.

* Role required: admin

### GET /api/users

Get all the users. Only used in the admin console.

* Role required: user

### GET /api/users/:id

Get a single user. Only used in the admin console.

* Role required: user

### PUT /api/users/:id

Update a user. Only used in the admin console.

* Role required: admin

### DELETE /api/users/:id

Delete a user. Only used in the admin console.

* Role required: admin