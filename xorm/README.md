Result:
```text
2022/11/28 21:48:13 ok: connected to database
2022/11/28 21:48:13 ok: episodes table is created
2022/11/28 21:48:13 ok: test/episodes table is created
2022/11/28 21:48:13 ok: series table is created
2022/11/28 21:48:13 ok: seasons table is created
2022/11/28 21:48:13 ok: fill tables with data
2022/11/28 21:48:13 Replace by fetch data
2022/11/28 21:48:14 > number of rows of `test/episodes`: 70
2022/11/28 21:48:14 > number of rows after replace: 70
2022/11/28 21:48:14 Explain Query
2022/11/28 21:48:14 AST (
(let $1 '('"/local/series" '"1" '"72075186232723360:1536"))
(let $2 '('"series_id" (Nothing (OptionalType (DataType 'String))) (Void)))
(let $3 '('"IncFrom" '"IncTo" $2))
(let $4 '('"release_date" '"series_id" '"title"))
(let $5 (Uint64 '"1001"))
(let $6 (KiSelectRange '"db" $1 $3 $4 '('('"ItemsLimit" $5))))
(return '('((Take $6 $5)) (List (ListType (VoidType)))))
)

2022/11/28 21:48:14 Plan {
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
2022/11/28 21:48:14 Select Default
2022/11/28 21:48:14 > select of all known series:
2022/11/28 21:48:14 > [46704bd4-ca83-49bd-bf59-98b1f70552b4] Silicon Valley (2014-04-06)
2022/11/28 21:48:14 > [6d8335b7-3e52-4f36-8022-cbee1777f103] IT Crowd (2006-02-03)
2022/11/28 21:48:14 Select Scan
2022/11/28 21:48:14 > scan select of episodes of `Season 1` of `IT Crowd` between 2006-01-01 and 2006-12-31:
2022/11/28 21:48:14 > [2d928dc5-8831-460a-8898-f9762d5c5e3a] Calamity Jen (2006-02-03)
2022/11/28 21:48:14 > [f19e2396-1496-443e-b5b3-71b3a9d706ea] Yesterday's Jam (2006-02-03)
2022/11/28 21:48:14 > [b1b95d9c-5227-4e77-8db6-d5a8ec175366] Fifty-Fifty (2006-02-10)
2022/11/28 21:48:14 > [04ddc73f-d5d9-4455-9d66-856412e965c3] The Red Door (2006-02-17)
2022/11/28 21:48:14 > [a4192b85-84f1-476b-b6de-eeafbda966f1] The Haunting of Bill Crouse (2006-02-24)
2022/11/28 21:48:14 > [4ec2e90f-d494-4277-bb70-ff94875048d2] Aunt Irma Visits (2006-03-03)
2022/11/28 21:48:14 INNER JOIN
2022/11/28 21:48:14 got 9 records
2022/11/28 21:48:14 > 16188f24-c81b-46db-ad32-79874d26e591  Season 1  46704bd4-ca83-49bd-bf59-98b1f70552b4  Silicon Valley  
2022/11/28 21:48:14 > 236611ac-f348-45bf-a997-ef6a171aa950  Season 4  46704bd4-ca83-49bd-bf59-98b1f70552b4  Silicon Valley  
2022/11/28 21:48:14 > 9c2134d0-bcca-43d0-bb8f-d6eecbcc0577  Season 3  46704bd4-ca83-49bd-bf59-98b1f70552b4  Silicon Valley  
2022/11/28 21:48:14 > aa49a385-6ec9-4fde-918a-d4c6d37675a7  Season 5  46704bd4-ca83-49bd-bf59-98b1f70552b4  Silicon Valley  
2022/11/28 21:48:14 > bea64e93-f937-4897-b37b-6b6d4934dbd7  Season 2  46704bd4-ca83-49bd-bf59-98b1f70552b4  Silicon Valley  
2022/11/28 21:48:14 > 46cd46fc-0343-41c6-8210-0559ba8a3781  Season 2  6d8335b7-3e52-4f36-8022-cbee1777f103  IT Crowd  
2022/11/28 21:48:14 > e301428e-f140-4c31-ac46-c2c182eacfc6  Season 4  6d8335b7-3e52-4f36-8022-cbee1777f103  IT Crowd  
2022/11/28 21:48:14 > e4ceb21f-8483-46bf-b098-c09a95c19877  Season 3  6d8335b7-3e52-4f36-8022-cbee1777f103  IT Crowd  
2022/11/28 21:48:14 > ea70d62f-64a7-40bc-aba9-b4aba0ae02af  Season 1  6d8335b7-3e52-4f36-8022-cbee1777f103  IT Crowd  
2022/11/28 21:48:14 LEFT JOIN
2022/11/28 21:48:14 got 9 records
2022/11/28 21:48:14 > 16188f24-c81b-46db-ad32-79874d26e591  Season 1  46704bd4-ca83-49bd-bf59-98b1f70552b4  Silicon Valley  
2022/11/28 21:48:14 > 236611ac-f348-45bf-a997-ef6a171aa950  Season 4  46704bd4-ca83-49bd-bf59-98b1f70552b4  Silicon Valley  
2022/11/28 21:48:14 > 9c2134d0-bcca-43d0-bb8f-d6eecbcc0577  Season 3  46704bd4-ca83-49bd-bf59-98b1f70552b4  Silicon Valley  
2022/11/28 21:48:14 > aa49a385-6ec9-4fde-918a-d4c6d37675a7  Season 5  46704bd4-ca83-49bd-bf59-98b1f70552b4  Silicon Valley  
2022/11/28 21:48:14 > bea64e93-f937-4897-b37b-6b6d4934dbd7  Season 2  46704bd4-ca83-49bd-bf59-98b1f70552b4  Silicon Valley  
2022/11/28 21:48:14 > 46cd46fc-0343-41c6-8210-0559ba8a3781  Season 2  6d8335b7-3e52-4f36-8022-cbee1777f103  IT Crowd  
2022/11/28 21:48:14 > e301428e-f140-4c31-ac46-c2c182eacfc6  Season 4  6d8335b7-3e52-4f36-8022-cbee1777f103  IT Crowd  
2022/11/28 21:48:14 > e4ceb21f-8483-46bf-b098-c09a95c19877  Season 3  6d8335b7-3e52-4f36-8022-cbee1777f103  IT Crowd  
2022/11/28 21:48:14 > ea70d62f-64a7-40bc-aba9-b4aba0ae02af  Season 1  6d8335b7-3e52-4f36-8022-cbee1777f103  IT Crowd  
2022/11/28 21:48:14 RIGHT JOIN
2022/11/28 21:48:14 got 9 records
2022/11/28 21:48:14 > 16188f24-c81b-46db-ad32-79874d26e591  Season 1  46704bd4-ca83-49bd-bf59-98b1f70552b4  Silicon Valley  
2022/11/28 21:48:14 > 236611ac-f348-45bf-a997-ef6a171aa950  Season 4  46704bd4-ca83-49bd-bf59-98b1f70552b4  Silicon Valley  
2022/11/28 21:48:14 > 9c2134d0-bcca-43d0-bb8f-d6eecbcc0577  Season 3  46704bd4-ca83-49bd-bf59-98b1f70552b4  Silicon Valley  
2022/11/28 21:48:14 > aa49a385-6ec9-4fde-918a-d4c6d37675a7  Season 5  46704bd4-ca83-49bd-bf59-98b1f70552b4  Silicon Valley  
2022/11/28 21:48:14 > bea64e93-f937-4897-b37b-6b6d4934dbd7  Season 2  46704bd4-ca83-49bd-bf59-98b1f70552b4  Silicon Valley  
2022/11/28 21:48:14 > 46cd46fc-0343-41c6-8210-0559ba8a3781  Season 2  6d8335b7-3e52-4f36-8022-cbee1777f103  IT Crowd  
2022/11/28 21:48:14 > e301428e-f140-4c31-ac46-c2c182eacfc6  Season 4  6d8335b7-3e52-4f36-8022-cbee1777f103  IT Crowd  
2022/11/28 21:48:14 > e4ceb21f-8483-46bf-b098-c09a95c19877  Season 3  6d8335b7-3e52-4f36-8022-cbee1777f103  IT Crowd  
2022/11/28 21:48:14 > ea70d62f-64a7-40bc-aba9-b4aba0ae02af  Season 1  6d8335b7-3e52-4f36-8022-cbee1777f103  IT Crowd  
2022/11/28 21:48:14 FULL JOIN
2022/11/28 21:48:14 got 9 records
2022/11/28 21:48:14 > 16188f24-c81b-46db-ad32-79874d26e591  Season 1  46704bd4-ca83-49bd-bf59-98b1f70552b4  Silicon Valley  
2022/11/28 21:48:14 > 236611ac-f348-45bf-a997-ef6a171aa950  Season 4  46704bd4-ca83-49bd-bf59-98b1f70552b4  Silicon Valley  
2022/11/28 21:48:14 > 9c2134d0-bcca-43d0-bb8f-d6eecbcc0577  Season 3  46704bd4-ca83-49bd-bf59-98b1f70552b4  Silicon Valley  
2022/11/28 21:48:14 > aa49a385-6ec9-4fde-918a-d4c6d37675a7  Season 5  46704bd4-ca83-49bd-bf59-98b1f70552b4  Silicon Valley  
2022/11/28 21:48:14 > bea64e93-f937-4897-b37b-6b6d4934dbd7  Season 2  46704bd4-ca83-49bd-bf59-98b1f70552b4  Silicon Valley  
2022/11/28 21:48:14 > 46cd46fc-0343-41c6-8210-0559ba8a3781  Season 2  6d8335b7-3e52-4f36-8022-cbee1777f103  IT Crowd  
2022/11/28 21:48:14 > e301428e-f140-4c31-ac46-c2c182eacfc6  Season 4  6d8335b7-3e52-4f36-8022-cbee1777f103  IT Crowd  
2022/11/28 21:48:14 > e4ceb21f-8483-46bf-b098-c09a95c19877  Season 3  6d8335b7-3e52-4f36-8022-cbee1777f103  IT Crowd  
2022/11/28 21:48:14 > ea70d62f-64a7-40bc-aba9-b4aba0ae02af  Season 1  6d8335b7-3e52-4f36-8022-cbee1777f103  IT Crowd  
2022/11/28 21:48:14 LEFT SEMI JOIN
2022/11/28 21:48:14 got 9 records
2022/11/28 21:48:14 > 16188f24-c81b-46db-ad32-79874d26e591  Season 1  
2022/11/28 21:48:14 > 236611ac-f348-45bf-a997-ef6a171aa950  Season 4  
2022/11/28 21:48:14 > 46cd46fc-0343-41c6-8210-0559ba8a3781  Season 2  
2022/11/28 21:48:14 > 9c2134d0-bcca-43d0-bb8f-d6eecbcc0577  Season 3  
2022/11/28 21:48:14 > aa49a385-6ec9-4fde-918a-d4c6d37675a7  Season 5  
2022/11/28 21:48:14 > bea64e93-f937-4897-b37b-6b6d4934dbd7  Season 2  
2022/11/28 21:48:14 > e301428e-f140-4c31-ac46-c2c182eacfc6  Season 4  
2022/11/28 21:48:14 > e4ceb21f-8483-46bf-b098-c09a95c19877  Season 3  
2022/11/28 21:48:14 > ea70d62f-64a7-40bc-aba9-b4aba0ae02af  Season 1  
2022/11/28 21:48:14 Update table
2022/11/28 21:48:14 ok: after episodes table is updated
2022/11/28 21:48:14 > test 999 2014-04-13
2022/11/28 21:48:14 > test 999 2017-05-21
2022/11/28 21:48:14 > test 999 2017-06-04
2022/11/28 21:48:14 > test 999 2017-06-11
2022/11/28 21:48:14 > test 999 2016-05-22
2022/11/28 21:48:14 > test 999 2016-06-26
2022/11/28 21:48:14 > test 999 2015-05-03
2022/11/28 21:48:14 Delete Records
2022/11/28 21:48:14 > before delete: 9
2022/11/28 21:48:14 ok: deleted all records with first_aired from 2007-06-01 to 2008-06-01 in all seasons
2022/11/28 21:48:14 > after delete: 7
```