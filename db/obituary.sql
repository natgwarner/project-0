/*This database will be used to track player playthroughs, 
showing either which ending of the game they got or where they died.*/

create table obituary (
    id integer primary KEY,
    playername VARCHAR NOT NULL,
    playerdeath VARCHAR NOT NULL
)