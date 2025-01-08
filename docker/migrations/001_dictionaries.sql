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

insert into flights (id, origin_id, destination_id, departure_datetime, arrival_datetime, company_id) values ('523f142f-dc19-4ee3-b138-21a0f7423266', '5398bb7d-8bcd-4542-ba97-3c60cee2259d', '72fd899d-610b-4d3b-b744-bd8667339c2f', TIMESTAMPTZ '2025-01-25 10:00:00 +03:00', TIMESTAMPTZ '2025-01-26 01:59:00 +09:00', '26f57140-bda6-48af-8077-684096171af1'), ( 'ce8efe9d-67c8-42d9-98eb-9416fd475a3b', '72fd899d-610b-4d3b-b744-bd8667339c2f', 'af937b22-ef36-4483-b679-303045969cee', TIMESTAMPTZ '2025-01-28 19:30:00 +09:00', TIMESTAMPTZ '2025-01-28 23:29:00 +03:00', '060e01b0-6695-4c6d-a9cc-6b8c25875fc2');

insert into class_data(flight_id, class_type, cost, luggage, buggage) values ('523f142f-dc19-4ee3-b138-21a0f7423266', 'E1', 1000, 10, 18), ('523f142f-dc19-4ee3-b138-21a0f7423266', 'E0', 800, 10, 0), ('523f142f-dc19-4ee3-b138-21a0f7423266', 'B0', 1599, 10, 23), ('ce8efe9d-67c8-42d9-98eb-9416fd475a3b', 'E1', 1251, 8, 23), ('ce8efe9d-67c8-42d9-98eb-9416fd475a3b', 'B2', 2550, 10, 25);