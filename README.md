# raisingSuperstars-prodigy-program
As part of our Prodigy Programs, we provide 5-minute per day plans to users, which they can view each day as a list and mark as complete in the app to track their progress.


As per my understanding there will 4 APIs.
1. To post the whole week plan into the Database.
    (There should be a cron which automatically runs every saturday or any day of the week, which will populate the db with the whole week plan).

    POST API - (/weekplan)

2. To show the "Suggested - Day wise", We will be having a get API.
    GET API- Get the week plan (/getweekplan?weekId=3)
    GET API- get the day plan (getweekplan?weekId=3&dayNumber=1)


3. To mark the progress for that particular day by the user, we will have 
    3.1- Post/update the progress of the day.(will accept array of activities from front end)
    UPDATE API - (/updatedayplan?dayNumber=1)

4. User registring API
    POST API- (/registerUser)

Schema: 
4 tables 
4. User table
    id, name, email

1. WeekPlan
    id, Desc

3. Activity Desc(master table)
    id, category, time, freq, activity.
    The data will be almost constant.

2. Day plan
    id, user_id(FK), weekId(FK), activityId(FK), day_number, completed_at, created_at, updated_at


The sql commands are present in the postgre.sql file.
The code is structured and working.

