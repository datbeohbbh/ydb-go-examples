Result
```text
2022/11/27 23:43:01 ok: connected to database
2022/11/27 23:43:01 ok: episodes table is created
2022/11/27 23:43:01 ok: series table is created
2022/11/27 23:43:02 ok: seasons table is created
2022/11/27 23:43:02 ok: fill tables with data
2022/11/27 23:43:02 Explain Query
2022/11/27 23:43:02 AST (
(let $1 '('"/local/series" '"1" '"72075186232723360:812"))
(let $2 '('"series_id" (Nothing (OptionalType (DataType 'String))) (Void)))
(let $3 '('"IncFrom" '"IncTo" $2))
(let $4 '('"release_date" '"series_id" '"title"))
(let $5 (Uint64 '"1001"))
(let $6 (KiSelectRange '"db" $1 $3 $4 '('('"ItemsLimit" $5))))
(return '('((Take $6 $5)) (List (ListType (VoidType)))))
)

2022/11/27 23:43:02 Plan {
  "meta":
    {
      "version":"0.1",
      "type":"query"
    },
  "tables":
    [
      {
        "name":"\/local\/series",
        "reads":
          [
            {
              "type":"FullScan",
              "scan_by":
                [
                  "series_id"
                ],
              "limit":"\"1001\"",
              "columns":
                [
                  "release_date",
                  "series_id",
                  "title"
                ]
            }
          ]
      }
    ]
}
2022/11/27 23:43:02 Select Default
2022/11/27 23:43:02 > select of all known series:
2022/11/27 23:43:02 > [13e282be-6aff-41e7-96f8-d2c88dd15354] Silicon Valley (2014-04-06)
2022/11/27 23:43:02 > [5da77983-9aae-45df-8dd0-4bf65086eb5b] IT Crowd (2006-02-03)
2022/11/27 23:43:02 Select Scan
2022/11/27 23:43:02 > scan select of episodes of `Season 1` of `IT Crowd` between 2006-01-01 and 2006-12-31:
2022/11/27 23:43:02 > [21806f39-0574-482f-9223-bbafab187826] Calamity Jen (2006-02-03)
2022/11/27 23:43:02 > [c07f5ef0-acad-4306-b7b7-f6b9b4dae92f] Yesterday's Jam (2006-02-03)
2022/11/27 23:43:02 > [4a5571f1-555e-4cdc-b02d-44a221d766b7] Fifty-Fifty (2006-02-10)
2022/11/27 23:43:02 > [75f17c1d-5f67-401b-b3f6-06bc42c13954] The Red Door (2006-02-17)
2022/11/27 23:43:02 > [9e68eb4c-d759-4f10-bda5-2f009a1a0908] The Haunting of Bill Crouse (2006-02-24)
2022/11/27 23:43:02 > [d61c2c15-45d8-4eff-b8a2-ce345cc50d74] Aunt Irma Visits (2006-03-03)
2022/11/27 23:43:02 INNER JOIN
2022/11/27 23:43:02 got 9 records
2022/11/27 23:43:02 > 3100edde-2cce-4703-82c9-110513bb146a  Season 4  13e282be-6aff-41e7-96f8-d2c88dd15354  Silicon Valley  
2022/11/27 23:43:02 > 477def1a-fd85-496d-a4b2-10f34e2463b2  Season 3  13e282be-6aff-41e7-96f8-d2c88dd15354  Silicon Valley  
2022/11/27 23:43:02 > 7f504e62-7f38-4306-820b-8300de01c328  Season 5  13e282be-6aff-41e7-96f8-d2c88dd15354  Silicon Valley  
2022/11/27 23:43:02 > c254dde8-c17d-4d10-a1ff-9e4b0a00b9df  Season 2  13e282be-6aff-41e7-96f8-d2c88dd15354  Silicon Valley  
2022/11/27 23:43:02 > dade1cf4-2236-498c-bc45-3bfc905b10f3  Season 1  13e282be-6aff-41e7-96f8-d2c88dd15354  Silicon Valley  
2022/11/27 23:43:02 > 1a173798-392e-4f68-9e5c-c15b65059dd9  Season 1  5da77983-9aae-45df-8dd0-4bf65086eb5b  IT Crowd  
2022/11/27 23:43:02 > 39e6e1d8-d662-42e8-a024-6162b89308a9  Season 4  5da77983-9aae-45df-8dd0-4bf65086eb5b  IT Crowd  
2022/11/27 23:43:02 > 6b9ba703-0a5b-4e15-be1a-27fdecde981a  Season 3  5da77983-9aae-45df-8dd0-4bf65086eb5b  IT Crowd  
2022/11/27 23:43:02 > f2731a96-aaf4-40df-b80b-0299a2a7dba9  Season 2  5da77983-9aae-45df-8dd0-4bf65086eb5b  IT Crowd  
2022/11/27 23:43:02 LEFT JOIN
2022/11/27 23:43:02 got 9 records
2022/11/27 23:43:02 > 3100edde-2cce-4703-82c9-110513bb146a  Season 4  13e282be-6aff-41e7-96f8-d2c88dd15354  Silicon Valley  
2022/11/27 23:43:02 > 477def1a-fd85-496d-a4b2-10f34e2463b2  Season 3  13e282be-6aff-41e7-96f8-d2c88dd15354  Silicon Valley  
2022/11/27 23:43:02 > 7f504e62-7f38-4306-820b-8300de01c328  Season 5  13e282be-6aff-41e7-96f8-d2c88dd15354  Silicon Valley  
2022/11/27 23:43:02 > c254dde8-c17d-4d10-a1ff-9e4b0a00b9df  Season 2  13e282be-6aff-41e7-96f8-d2c88dd15354  Silicon Valley  
2022/11/27 23:43:02 > dade1cf4-2236-498c-bc45-3bfc905b10f3  Season 1  13e282be-6aff-41e7-96f8-d2c88dd15354  Silicon Valley  
2022/11/27 23:43:02 > 1a173798-392e-4f68-9e5c-c15b65059dd9  Season 1  5da77983-9aae-45df-8dd0-4bf65086eb5b  IT Crowd  
2022/11/27 23:43:02 > 39e6e1d8-d662-42e8-a024-6162b89308a9  Season 4  5da77983-9aae-45df-8dd0-4bf65086eb5b  IT Crowd  
2022/11/27 23:43:02 > 6b9ba703-0a5b-4e15-be1a-27fdecde981a  Season 3  5da77983-9aae-45df-8dd0-4bf65086eb5b  IT Crowd  
2022/11/27 23:43:02 > f2731a96-aaf4-40df-b80b-0299a2a7dba9  Season 2  5da77983-9aae-45df-8dd0-4bf65086eb5b  IT Crowd  
2022/11/27 23:43:02 RIGHT JOIN
2022/11/27 23:43:02 got 9 records
2022/11/27 23:43:02 > 3100edde-2cce-4703-82c9-110513bb146a  Season 4  13e282be-6aff-41e7-96f8-d2c88dd15354  Silicon Valley  
2022/11/27 23:43:02 > 477def1a-fd85-496d-a4b2-10f34e2463b2  Season 3  13e282be-6aff-41e7-96f8-d2c88dd15354  Silicon Valley  
2022/11/27 23:43:02 > 7f504e62-7f38-4306-820b-8300de01c328  Season 5  13e282be-6aff-41e7-96f8-d2c88dd15354  Silicon Valley  
2022/11/27 23:43:02 > c254dde8-c17d-4d10-a1ff-9e4b0a00b9df  Season 2  13e282be-6aff-41e7-96f8-d2c88dd15354  Silicon Valley  
2022/11/27 23:43:02 > dade1cf4-2236-498c-bc45-3bfc905b10f3  Season 1  13e282be-6aff-41e7-96f8-d2c88dd15354  Silicon Valley  
2022/11/27 23:43:02 > 1a173798-392e-4f68-9e5c-c15b65059dd9  Season 1  5da77983-9aae-45df-8dd0-4bf65086eb5b  IT Crowd  
2022/11/27 23:43:02 > 39e6e1d8-d662-42e8-a024-6162b89308a9  Season 4  5da77983-9aae-45df-8dd0-4bf65086eb5b  IT Crowd  
2022/11/27 23:43:02 > 6b9ba703-0a5b-4e15-be1a-27fdecde981a  Season 3  5da77983-9aae-45df-8dd0-4bf65086eb5b  IT Crowd  
2022/11/27 23:43:02 > f2731a96-aaf4-40df-b80b-0299a2a7dba9  Season 2  5da77983-9aae-45df-8dd0-4bf65086eb5b  IT Crowd  
2022/11/27 23:43:02 FULL JOIN
2022/11/27 23:43:02 got 9 records
2022/11/27 23:43:02 > 3100edde-2cce-4703-82c9-110513bb146a  Season 4  13e282be-6aff-41e7-96f8-d2c88dd15354  Silicon Valley  
2022/11/27 23:43:02 > 477def1a-fd85-496d-a4b2-10f34e2463b2  Season 3  13e282be-6aff-41e7-96f8-d2c88dd15354  Silicon Valley  
2022/11/27 23:43:02 > 7f504e62-7f38-4306-820b-8300de01c328  Season 5  13e282be-6aff-41e7-96f8-d2c88dd15354  Silicon Valley  
2022/11/27 23:43:02 > c254dde8-c17d-4d10-a1ff-9e4b0a00b9df  Season 2  13e282be-6aff-41e7-96f8-d2c88dd15354  Silicon Valley  
2022/11/27 23:43:02 > dade1cf4-2236-498c-bc45-3bfc905b10f3  Season 1  13e282be-6aff-41e7-96f8-d2c88dd15354  Silicon Valley  
2022/11/27 23:43:02 > 1a173798-392e-4f68-9e5c-c15b65059dd9  Season 1  5da77983-9aae-45df-8dd0-4bf65086eb5b  IT Crowd  
2022/11/27 23:43:02 > 39e6e1d8-d662-42e8-a024-6162b89308a9  Season 4  5da77983-9aae-45df-8dd0-4bf65086eb5b  IT Crowd  
2022/11/27 23:43:02 > 6b9ba703-0a5b-4e15-be1a-27fdecde981a  Season 3  5da77983-9aae-45df-8dd0-4bf65086eb5b  IT Crowd  
2022/11/27 23:43:02 > f2731a96-aaf4-40df-b80b-0299a2a7dba9  Season 2  5da77983-9aae-45df-8dd0-4bf65086eb5b  IT Crowd  
2022/11/27 23:43:02 LEFT SEMI JOIN
2022/11/27 23:43:02 got 9 records
2022/11/27 23:43:02 > 1a173798-392e-4f68-9e5c-c15b65059dd9  Season 1  
2022/11/27 23:43:02 > 3100edde-2cce-4703-82c9-110513bb146a  Season 4  
2022/11/27 23:43:02 > 39e6e1d8-d662-42e8-a024-6162b89308a9  Season 4  
2022/11/27 23:43:02 > 477def1a-fd85-496d-a4b2-10f34e2463b2  Season 3  
2022/11/27 23:43:02 > 6b9ba703-0a5b-4e15-be1a-27fdecde981a  Season 3  
2022/11/27 23:43:02 > 7f504e62-7f38-4306-820b-8300de01c328  Season 5  
2022/11/27 23:43:02 > c254dde8-c17d-4d10-a1ff-9e4b0a00b9df  Season 2  
2022/11/27 23:43:02 > dade1cf4-2236-498c-bc45-3bfc905b10f3  Season 1  
2022/11/27 23:43:02 > f2731a96-aaf4-40df-b80b-0299a2a7dba9  Season 2  
2022/11/27 23:43:02 Update table
2022/11/27 23:43:02 ok: after episodes table is updated
2022/11/27 23:43:02 > test 999 2017-06-11
2022/11/27 23:43:02 > test 999 2017-05-21
2022/11/27 23:43:02 > test 999 2017-06-04
2022/11/27 23:43:02 > test 999 2016-06-26
2022/11/27 23:43:02 > test 999 2016-05-22
2022/11/27 23:43:02 > test 999 2015-05-03
2022/11/27 23:43:02 > test 999 2014-04-13
2022/11/27 23:43:02 Delete Records
2022/11/27 23:43:02 > before delete: 9
2022/11/27 23:43:02 ok: deleted all records with first_aired from 2007-06-01 to 2008-06-01 in all seasons
2022/11/27 23:43:02 > after delete: 7
```