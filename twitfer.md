# Twitfer

<hr>

## Requirement Gathering

### Functional Requirement
- [ ] User timeline: Displaying user’s tweets and tweets user retweet
- [ ] Home timeline: Displaying Tweets from people user follow
- [ ] Search timeline: Display search results based on #tags or keyword
- [ ] The user should be able to follow another user
- [ ] The user should see trends

### Non-functional Requirement
- [ ] The user should be able to tweet as fast as possible
- [ ] Users should be able to tweet millions of followers within a few seconds (5 seconds)

<hr>

## Consideration

- Read requests >> Write requests
- High Availibility > Consistency.

<hr>

## Building

### Database Postgresql

User Table & Tweet Table
- User Table = (userid, username, password, email, etc)
- Tweet Table = (userid, tweetid, datetime)
- User Table will have 1 to many relationships with Tweet Table

Follower Table
- Followers Table = (userid, followerid)
- When a user follows another user, it gets stored in Followers Table, and also cache it Redis

### Caching Redis
- **userid-tweets** = [tweetid_1, tweetid_2, ..., tweetid_n]
- **userid-followers** = [userid_1, userid2, ...., userid_n]

### User Timeline(userid == idx)
-  from Redis GET idx-tweets ORDER BY datetime
-  if cache-missed, get from Postgresql tweet table

### Home Timeline idy
- user(userid = idx) tweet twx
- Tweet Table.insert(idx, twx)
- send tweet to all **idx-followers** Home Timeline Rabbitmq Queue, one of it is **idy**
- get all celebrities that **idy** following
- merge celebrities queue and rakjel queue 

### Trending - Storm Kafka Stream
- user(userid = idx) tweet twx
- Tweet Table.insert(idx, twx)
- get all hastags
- filter offensive hastags
- for each h in hastags: INCR redis
- TODO: specify

### Search
- if @* ==> search username
- if not ==> search tweet
- 