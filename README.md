# InstagramBackenedAPI


### 1. Created Member1 User in the CREATE USER

- Input
```
{
    "name":"Mamber1",
    "email":"member1@gmail.com",
    "password":"sj@78399k24"
}
```

### 2. Input the request CREATE USER: `http://localhost:9000/users`
![CreateUserId](https://user-images.githubusercontent.com/67019423/136685574-6ffb427e-2b27-41c6-94c3-058fe97ca7e9.png)


### 3. received ID: 6162846c0d336b6e985185c9

```
{
    "id": "6162846c0d336b6e985185c9",
    "name": "Mamber1",
    "email": "member1@gmail.com",
    "password": "sj@78399k24"
}
```

### 4. Input the request in the GET USER ID: `http://localhost:9000/users/6162846c0d336b6e985185c9`

### 5: Got the data from the requested id:
```
{
    "id": "6162846c0d336b6e985185c9",
    "name": "Mamber1",
    "email": "member1@gmail.com",
    "password": "sj@78399k24"
}
```

![GetPost](https://user-images.githubusercontent.com/67019423/136685582-6f938e7d-2d68-4b06-839d-adc900816ec5.png)

### 6. Database Created by requesting all the data from the users's server from the postman api

![usersDatabase](https://user-images.githubusercontent.com/67019423/136685589-fb3a0a6f-77aa-40a1-97f2-02ec68c8adaa.png)


### 7. Created ID for Member1 User's Post in the CREATE POST
### 8. Input the request CREATE USER: `http://localhost:9000/posts`

- Input
```
{
    "caption":"Caption4 Hey there!",
    "image":"https://www.callicoder.com/static/87b2ad19dce949b603798cf6b1fd1ab6/e5166/golang-structs.jpg",
    "userid":"6162846c0d336b6e985185c9"
}
```

![CreatePost](https://user-images.githubusercontent.com/67019423/136685590-97602370-9e1f-4fa0-870d-f76c15f09a62.png)

### 9. received ID: 616286030d336b3fdc55db58
```
{
    "id": "616286030d336b3fdc55db58",
    "caption": "Caption4 Hey there!",
    "image": "https://www.callicoder.com/static/87b2ad19dce949b603798cf6b1fd1ab6/e5166/golang-structs.jpg",
    "userid": "6162846c0d336b6e985185c9",
    "time": "0001-01-01T00:00:00Z"
}
```

### 10. Input the request in the GET USER ID: `http://localhost:9000/posts/616286030d336b3fdc55db58`

### 11: Got the data from the requested post:
```
{
    "id": "616286030d336b3fdc55db58",
    "caption": "Caption4 Hey there!",
    "image": "https://www.callicoder.com/static/87b2ad19dce949b603798cf6b1fd1ab6/e5166/golang-structs.jpg",
    "userid": "6162846c0d336b6e985185c9",
    "time": "0001-01-01T00:00:00Z"
}
```

![GetUserId](https://user-images.githubusercontent.com/67019423/136685587-257228c1-e216-4f6d-a762-49bcf4840ca5.png)

### 12. Database Created by requesting all the data from the posts's server from the postman api

![postsDatabase](https://user-images.githubusercontent.com/67019423/136685588-705245ab-55d3-4f3c-bc25-7e787518228f.png)
