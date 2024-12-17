```
CREATE TABLE IF NOT EXISTS public."My_game"
(
    id serial NOT NULL,
    question text,
    ans1 character varying(50),
    ans2 character varying(50),
    ans3 character varying(50),
    ans4 character varying(50),
    points integer,
    CONSTRAINT "My_game_pkey" PRIMARY KEY (id)
);
```


//get - на страницу вопрос 
//post(id, number_ans) -> true/false
//getscore 