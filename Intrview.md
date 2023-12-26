1. **Introduction (2-3 minutes):**
   - Introduction
     - Name
     - Native Place
     - Current Position
     - Current Company

2. **Summary of Experience (3-4 minutes):**
   - Highlight your 4+ years of experience as a Backend Developer.
      - So i have worked on many technologies during this period.
      - Transition from GeekyAnts to Glue Labs
      - Leading the backend team at GlueLabs
      - Project management, designing architecture of the application & scaling a particular feature
      - Go to Guy for budiling POCs & analysing the pros & cons of the system.
   - Mention your specialization in Node.js, Golang, and Python.
     -  it includes Node.Js, Golang, Php, Python
        - Node.Js
          - I'm proficient in Node.js & its corresponding frameworks including Express.Js, Nest.Js
        - Golang
          - Have 2 years of experience in Golang. Worked on a mico-service featuring Golang
          - Have built a open-source boilerplate repo using Gin Framework.
        - PHP
          - Worked on Laravel framework
        - Python
          - Have on & off experience. Worked on some projects including Python
          - Have worked on Django framwork
        - Experience across various domains
          - domain
            - healthcare
            - fintech
            - social media
   - Emphasize your achievements, such as optimizing a real-time audio conference app, constructing a scalable notification system, and implementing a central authentication system at GlueLabs.
     - For real-time audio conference app
       - Refactoring the code
       - The delete login was wrong, instead of bulk delete, we used to do it sequencially
       - The time to join the room was very high. Refactoring & applying some tweaks to the existing logic improved it a lot
       - Handling heavy track
       - Moved some extra data to redis to avoid latency incurred by DB call.
     - Notification System
       - So intially the notification system was fast but not scalable, we used to update the notification for each user in run-time, calculating & updating the rows frequenctly
       - Leveraged Kafka to sent every activity
       - At the other end, had a consumer
         - We dumped the data in DB directly
         - At runtime, we crunched the entries based on the notification scheme we had using a complex SQL & cached them for the time being.
     - Followed OpenId protoocl
     - Feed feature
       - Everything was coming from DB.
       - We moved the most popular data to redis.
       - Leveraged SortedSets for this
       - Sorted sorts only kept ids of posts based on the timestamp.
       - kep the data separately as a string

3. **Experience at GlueLabs (5-6 minutes):**
   - Provide more details about your current role at GlueLabs, focusing on the key projects and technologies you've worked with.
     - A senior software engineer at GlueLabs,
     - Working on 3 products internally.
     - One project FIFO is live
   - Discuss how you optimized the real-time audio conference app, detailing the tech stack used (Node.js, Websockets, Redis) and the impact on latency.
   - Explain your role in constructing the scalable notification system using Kafka and the central authentication system using Node.js and Next.js.
   - Mention your leadership role, including orchestrating project management, leading the backend team, and spearheading the architectural design of key features.

4. **Experience at GeekyAnts (4-5 minutes):**
   - Discuss your tenure at GeekyAnts, emphasizing your contributions to backend solutions using Node.js, Golang, PHP, and Python.
   - Highlight your involvement in a fintech project, stabilization of crashing environments, and contributions to multiple microservices.
     - Intiially a junio developer.
     - Environemnts were always crashing, Due to some dependant micro-services failing. 
     - Had to debug & fix the issues. After a month, eveything was stabilized.
     - Created new APIs
       - Used Loopback
       - PostgresQL
       - DbLink - a thing I explored while solving some issuess related to joins.
     - Gained deep understanding of micro-services
   - Mention your progression to a leadership role, leading a team, and participating in meetings focused on project architecture and workflow optimization.
     - Initially i was there as a junior developer, working on carshing pods
     - Then I got involved in more features, like
       - Rebalancing, 
       - KYC cron not being efficient
       - Liquidation of funds
       - Calculating PnL
       - Micro Service Architecture
       - Lead a team while handling multiple projects
     - I knew about eco-system & how things were working

5. **Skills and Technologies (2-3 minutes):**
   - Briefly outline your skills, including frontend (React.js, Next.js) and backend (Node.js, Express.js, Nest.js, Golang, Python) technologies.
     - Node.Js
       - a community based learning platform
       - Fintech projects using framework Loopback
       - a product we are building right now
   - Highlight your proficiency in databases (MySQL, MariaDB, PostgreSQL, MongoDB) and cloud platforms (AWS services, Kubernetes, Kafka, RabbitMQ).
     - MySQL
       - Initially started on some projects using MySQL
       - Built a 
         - company portal, using Laravel, MySQL
         -  an all-in-one social, communication and digital life management app designed to connect family, friends and community in a safe and secure space using MySQL, Laravel
     - PostgresQL
       - Started on Postgres after working on fintech project
       - Used procedures, triggers, materialized views
       - Worked on an open-source boilerplate repo that was built using Express.Js and PostgreSQL
         - Leveraged OOPs concepts of PostgreSQL
       - Complex JSON column manipulation queries
     - MongoDB
       - Worked on a taxi app that leveraged mongoDB to store its data
       - Like co-ordinates & all

6. **Projects (2-3 minutes):**
   - Discuss your significant projects, such as the Video Streaming Service and the AWS Guide.
     - Video Streaming Service
       - It started when I learnt about RTSP protocol.
       - I had some ip cameras with me.
       - Built a system that captures streams of these ip cameras, leveraged ffmpeg - one of the greatest tool I have encountered.
       - Transfer streams to another server, which will do further processing
         - Separeted audio & video, to reduce noise from audio & sticking it back.
         - Streamed it using HLS
   - Provide insights into the challenges faced and how you overcame them, showcasing your problem-solving skills
     - Streaming so many streams at the same time
       - Solution was to compress the streams initially & use a bigger infrastructure platform wiht higher bandiwth of I/Ops
     - Handling the streams that are transferred
       - I was using Node.Js, since it is single threaded, it becomes dificult to take adavantage of multi core architture
       - but Nodejs had an option of creating worker threads, so I tried to ulitize that.
     - Keeping audio & video getting lost
       - It was a tradeoff, intiially we used UDP for transfer of streams but more packets were getting lost,
       - We could use TCP, the packet loss would be less but now we added more overhead of acknowleding the stream, TCP adds network congestion

7. **Education (1-2 minutes):**
   - Mention your academic background, completing a B.E. in Computer Science at Chitkara University with a notable GPA.
   - Highlight any relevant coursework or projects during your education.

8. **Closing and Future Outlook (1-2 minutes):**
   1. Really into desinging scalable sytems, solving complex problems.
   2. My Journey has shaped me into a better developer at each point.
   3. Big fan of kuernetes, Golang by Google


Remember to tailor this structure to your comfort and the flow of the conversation. Good luck with your interview!