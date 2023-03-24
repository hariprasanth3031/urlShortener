# urlShortener System Design code in Go

# Tech Stack used
Backend programming - Golang, 
Database - MySql

# Functionalities
# Encode
Given a long url, converts into a tiny url based on a bijective function.
To create unique tinyurl's, we are going to take the autoincrement id generated in db and hash it using our own hash algorithm.

# Custom Hash Algorithm
We are going to consider character set [a-zA-Z0-9]. There are totally 62 characters. So for every id which is in base 10, we are going to onvert into base 64 and map it according to the character set.
Example: Character set mp= "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",  longurl = "google.com" and it inserted in the row 305.
Now, we are going to hash 122. 122 base 62 = [1,60]. Now map this in the character set. mp[1] = b, mp[60] = 8. So, the tiny url will be b8. In this way we are going to generate unique short url's as the id will be unique.

# Handling expiration
The generated short url needs to be expired in 24 hours form the point pf creation. we are going to store expiry time in db by adding one day to the time of creation of short url.

# Decode
Given a short url, fetch its corresponding longurl.
As we know, all the short url's are unique we can direct query in our db using the given tiny url and expire time >= current time.
Expire time >= current time condition will handle the expiration case.

# Scalabality on number of url's
As the character set consists of 62 characters, we can generate trillions of url.
In case we have fixed number of characters, n power k will be unique urls where n is length, k = 62



