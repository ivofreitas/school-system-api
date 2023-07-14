# SCHOOL-SYSTEM-API

## Requirements
*  Go 1.19 or higher. If you don't have Golang installed, you can download it from https://go.dev/doc/install)
*  Docker Compose (https://docs.docker.com/compose/install/)

## Setup

1. Install the project dependencies:

   ```
   make install
   ```

2. Configure the project:

   ```
   make configure
   ```

   Then edit `config/.env` with your desired configuration values.
   

3. Start local environment:

   ```
   make local-up
   ```

4. Start the project:

   ```
   make start
   ```

   The project should now be running at http://localhost:8088.


5. Run unit tests:

   ```
   make test
   ```
   
6. Run test coverage:

   ```
   make cover
   ```

## Usage

### Endpoints

The API has the following endpoints:

#### User

Create a user.
Required fields:
* username
* password
* role (teacher, student, admin)

##### Request

```
POST /v1/user
Content-Type: application/json

{
    "username": "john.doe",
    "password": "123456",
    "role": "teacher"
 }
```

##### Response

```
HTTP/1.1 201 Created
Content-Type: application/json

{
    "meta": {
        "offset": 0,
        "limit": 0,
        "record_count": 1
    },
    "records": [
        {
            "id": "125edfd8-a7bd-4ddd-8429-278e935def82",
            "username": "john.doe",
            "role": "teacher"
        }
    ]
}

```
------------------
#### Login

Login with a user.
Required fields:
* username
* password

##### Request

```
POST /v1/user/login
Content-Type: application/json

{
    "username": "john.doe",
    "password": "123456",
}
```

##### Response

```
HTTP/1.1 201 Created
Content-Type: application/json

{
    "meta": {
        "offset": 0,
        "limit": 0,
        "record_count": 1
    },
    "records": [
        {
            "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4iLCJ1c2VyIjoiam9obi5kb2UifQ.cqDH000_9wpwtp2pmrAgUmcPvlyNDxObz8ks6ohBiUU"
        }
    ]
}
```

------------------

#### Student

Create a student.

##### Request

```
POST /v1/student
Content-Type: application/json
Authorization: Bearer <token>

{
    "first_name": "joe",
    "last_name": "doe",
    "ssid": "1234"
}
```

##### Response

```
HTTP/1.1 201 Created
Content-Type: application/json

{
    "meta": {
        "offset": 0,
        "limit": 0,
        "record_count": 1
    },
    "records": [
        {
            "id": "5d03c45c-212b-4033-915a-fdfede91a0e9",
            "first_name": "joe",
            "last_name": "doe",
            "ssid": "1234",
            "created_at": "2023-07-13T23:56:02.679277-03:00",
            "updated_at": "2023-07-13T23:56:02.679277-03:00"
        }
    ]
}
```

------------------

Edit a student.

##### Request

```
PUT /v1/student/{student-id}
Content-Type: application/json
Authorization: Bearer <token>

{
    "first_name": "joe",
    "last_name": "doe",
    "ssid": "1234"
}
```

##### Response

```
HTTP/1.1 200 OK
Content-Type: application/json

{
    "meta": {
        "offset": 0,
        "limit": 0,
        "record_count": 1
    },
    "records": [
        {
            "id": "5d03c45c-212b-4033-915a-fdfede91a0e9",
            "first_name": "joe",
            "last_name": "doe",
            "ssid": "1234",
            "created_at": "2023-07-13T23:56:02.679277-03:00",
            "updated_at": "2023-07-13T23:56:02.679277-03:00"
        }
    ]
}
```

------------------

Delete a student.

##### Request

```
DELETE /v1/student/{student-id}
Authorization: Bearer <token>
```

##### Response

```
HTTP/1.1 204 No Content
```

------------------

#### Course

Create a course.

##### Request

```
POST /v1/course
Content-Type: application/json
Authorization: Bearer <token>

{
    "room_id": "0001",
    "max_students_number": 10,
    "course_name": "statistics"
}
```

##### Response

```
HTTP/1.1 201 Created
Content-Type: application/json

{
"meta": {
    "offset": 0,
    "limit": 0,
    "record_count": 1
},
"records": [
        {
            "id": "7164d890-79de-4bf3-a9aa-27b6d570c10b",
            "room_id": "0001",
            "max_students_number": 10,
            "course_name": "statistics",
            "created_at": "2023-07-13T23:55:41.554741-03:00",
            "updated_at": "2023-07-13T23:55:41.554741-03:00"
        }
    ]
}
```

------------------

Enroll students.

##### Request

```
POST /v1/course/student
Content-Type: application/json
Authorization: Bearer <token>

{
    "students": ["5d03c45c-212b-4033-915a-fdfede91a0e9"],
    "room_id": "0001"
}
```

##### Response

```
HTTP/1.1 201 Created
Content-Type: application/json

{
"meta": {
    "offset": 0,
    "limit": 0,
    "record_count": 1
},
"records": [
    {
        "students": [
            "5d03c45c-212b-4033-915a-fdfede91a0e9"
        ],
        "room_id": "0001"
    }
]
}
```

------------------

### Error Handling

If an error occurs, the API will return a JSON object with an error message:

```
{
    "developer_message": string,
    "user_message": string,
    "status_code": int
}
```

Possible HTTP status codes for errors include:

- `400 Bad Request` for invalid request data
- `401 Unauthorized` missing a valid authorization token
- `404 Not Found` resource not found
- `409 Conflict` username already present in the database
- `500 Internal Server Error` for server-side errors

## Swagger API Documentation

This project uses Swagger for API documentation. Swagger provides a user-friendly interface for exploring and testing the API.

To access the Swagger page:

1. Start the application if it's not already running.
2. Open a web browser and navigate to `http://localhost:8088/v1/swagger/index.html#/`.
3. The Swagger page should load, displaying a list of available endpoints.

From here, you can explore the available endpoints, see what parameters they require, and test them out.

If you have any questions or issues with the Swagger page, please refer to the API documentation or contact the project maintainers.