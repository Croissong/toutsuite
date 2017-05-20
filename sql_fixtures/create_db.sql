create table toutsuite (
        id integer not null primary key autoincrement,
        title text unique default 'dsa',
        url text unique,
        tags text default ''
        );
