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
		"role": "user"
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
		"role": "user"
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

Get a single review. 

Example response:

```
{
	"id": 35,
	"uuid": "31767b5a-3da6-49b6-a0a2-b3022423cbc6",
	"reviewer": {
		"id": 18,
		"username": "arjo.hooimeijer",
		"firstName": "Arjo",
		"surname": "Hooimeijer",
		"fullName": "Arjo Hooimeijer",
		"jobTitle": "Solution Architect ",
		"department": "Technology ",
		"company": "IceMobile",
		"email": "arjo.hooimeijer@icemobile.com",
		"avatar": "arjo.hooimeijer.png",
		"role": "user"
	},
	"reviewee": {
		"id": 14,
		"username": "rosa.vancolmjon",
		"firstName": "Rosa",
		"surname": "van Colmjon",
		"fullName": "Rosa van Colmjon",
		"jobTitle": "SCRUM Master",
		"department": "Project Management",
		"company": "ICEMOBILE",
		"email": "rosa.vancolmjon@icemobile.com",
		"avatar": "rosa.vancolmjon.png",
		"role": "user"
	},
	"cards": null,
	"remark": "Lorem ipsum",
	"completed": false
}
```

* Role required: user

### PUT /api/reviews/:id

Complete a review.

You only need to send the identifiers of the chosen cards, and the remark (to be extended to include one remark per card).

Example request:

```
PUT /api/reviews/36
{
	"cards": [{
		"id": 5
	}, {
		"id": 6
	}, {
		"id": 7
	}],
	"remark": "This guy smells like Cheeto's"
}
```

Example response:

```
{
	"id": 36,
	"uuid": "9fb11d44-41f2-431b-9455-bf1cb1ef2044",
	"reviewer": {
		"id": 18,
		"username": "arjo.hooimeijer",
		"firstName": "Arjo",
		"surname": "Hooimeijer",
		"fullName": "Arjo Hooimeijer",
		"jobTitle": "Solution Architect ",
		"department": "Technology ",
		"company": "IceMobile",
		"email": "arjo.hooimeijer@icemobile.com",
		"avatar": "arjo.hooimeijer.png",
		"role": "user"
	},
	"reviewee": {
		"id": 15,
		"username": "tanja.hattink",
		"firstName": "Tanja",
		"surname": "Hattink",
		"fullName": "Tanja Hattink",
		"jobTitle": "Head of People Operations",
		"department": "Human Resource",
		"company": "ICEMOBILE",
		"email": "tanja.hattink@icemobile.com",
		"avatar": "tanja.hattink.png",
		"role": "user"
	},
	"cards": [{
		"id": 5,
		"title": "Humorous",
		"category": 0,
		"epic": {
			"id": 3,
			"title": "Leadership"
		}
	}, {
		"id": 6,
		"title": "Good communicator",
		"category": 0,
		"epic": {
			"id": 3,
			"title": "Leadership"
		}
	}, {
		"id": 7,
		"title": "Creative",
		"category": 0,
		"epic": {
			"id": 3,
			"title": "Leadership"
		}
	}],
	"remark": "Ewww! This guy smells like Cheeto's :(",
	"completed": true
}
```

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
	"category": 0, // 0: Positive
	"epic": {
      	"id": 4,
      	"title": "Teamwork"
    }
}, {
	"id": 39,
	"title": "Always talking",
	"category": 1, // 1: Negative
	"epic": {
      	"id": 3,
      	"title": "Leadership"
    }
}, {
	"id": 16,
	"title": "Ambitious",
	"category": 0,
	"epic": {
      	"id": 3,
      	"title": "Leadership"
    }
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