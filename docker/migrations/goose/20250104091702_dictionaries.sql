-- +goose Up
-- +goose StatementBegin
insert into
    companies (id, name)
values (
        '060e01b0-6695-4c6d-a9cc-6b8c25875fc2',
        'Nordwind Airlines'
    ),
    (
        '26f57140-bda6-48af-8077-684096171af1',
        'RedWings'
    );

insert into
    countries (id, name)
values (
        '56a5f823-6a87-4561-9ec4-a94b9f563435',
        'Russia'
    ),
    (
        '15003160-1428-4e98-bdd1-839aaf29171f',
        'Japan'
    );

insert into
    cities (id, name, country_id)
values (
        '70dca730-0ed8-43f2-be7a-a72e1b25a4ec',
        'Moscow',
        '56a5f823-6a87-4561-9ec4-a94b9f563435'
    ),
    (
        'a39819bd-655f-4efa-8ef4-d9780e4280bd',
        'Tokyo',
        '15003160-1428-4e98-bdd1-839aaf29171f'
    );

insert into
    airports (id, airport_code, city_id)
values (
        '5398bb7d-8bcd-4542-ba97-3c60cee2259d',
        'DME',
        '70dca730-0ed8-43f2-be7a-a72e1b25a4ec'
    ),
    (
        'af937b22-ef36-4483-b679-303045969cee',
        'SVO',
        '70dca730-0ed8-43f2-be7a-a72e1b25a4ec'
    ),
    (
        '5ec8b1a6-abe6-4472-896b-840476ef19d1',
        'VKO',
        '70dca730-0ed8-43f2-be7a-a72e1b25a4ec'
    ),
    (
        '72fd899d-610b-4d3b-b744-bd8667339c2f',
        'HND',
        'a39819bd-655f-4efa-8ef4-d9780e4280bd'
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM airports;
DELETE FROM cities;
DELETE FROM countries;
DELETE FROM companies; 
-- +goose StatementEnd
