Result:
```text
2022/11/28 11:00:53 ok: connected to database
2022/11/28 11:00:53 ok: series table is created
2022/11/28 11:00:53 ok: seasons table is created
2022/11/28 11:00:53 ok: episodes table is created
2022/11/28 11:00:53 ok: fill tables with data
2022/11/28 11:00:53 Explain Query
2022/11/28 11:00:53 AST (
(let $1 '('"/local/series" '"1" '"72075186232723360:1161"))
(let $2 '('"series_id" (Nothing (OptionalType (DataType 'String))) (Void)))
(let $3 '('"IncFrom" '"IncTo" $2))
(let $4 '('"release_date" '"series_id" '"title"))
(let $5 (Uint64 '"1001"))
(let $6 (KiSelectRange '"db" $1 $3 $4 '('('"ItemsLimit" $5))))
(return '('((Take $6 $5)) (List (ListType (VoidType)))))
)

2022/11/28 11:00:53 Plan {
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
2022/11/28 11:00:53 Select Default
2022/11/28 11:00:53 > select of all known series:
2022/11/28 11:00:53 > [2655dfa4-65cf-463e-af3f-7293ba6fd198] Silicon Valley (2014-04-06)
2022/11/28 11:00:53 > [67c67307-a6af-4dd7-909d-cca3952a2373] IT Crowd (2006-02-03)
2022/11/28 11:00:53 Select Scan
2022/11/28 11:00:53 > scan select of episodes of `Season 1` of `IT Crowd` between 2006-01-01 and 2006-12-31:
2022/11/28 11:00:53 > [8543dcfd-c1d8-4c88-a4a4-fd66cc5e91ed] Calamity Jen (2006-02-03)
2022/11/28 11:00:53 > [0dbbb02c-0650-4424-869a-7cd3f1983e64] Yesterday's Jam (2006-02-03)
2022/11/28 11:00:53 > [51bd399d-59d6-499f-81d9-193d54fa52d5] Fifty-Fifty (2006-02-10)
2022/11/28 11:00:53 > [2c3cbd1b-fb2c-4c42-97d3-74d62c3a4e79] The Red Door (2006-02-17)
2022/11/28 11:00:53 > [4d9fee03-85b6-4b70-be48-65149e40206c] The Haunting of Bill Crouse (2006-02-24)
2022/11/28 11:00:53 > [2d133ead-16a6-4b64-8887-1f8420eada5f] Aunt Irma Visits (2006-03-03)
2022/11/28 11:00:53 INNER JOIN
2022/11/28 11:00:53 got 9 records
2022/11/28 11:00:53 > 16e5372f-f467-489e-8cc1-b16867befa6e  Season 5  2655dfa4-65cf-463e-af3f-7293ba6fd198  Silicon Valley  
2022/11/28 11:00:53 > 294701e8-9d72-4d87-8443-ca7f12985dc6  Season 3  2655dfa4-65cf-463e-af3f-7293ba6fd198  Silicon Valley  
2022/11/28 11:00:53 > 65f1547c-aac6-4ce9-af02-ca9cd809c893  Season 2  2655dfa4-65cf-463e-af3f-7293ba6fd198  Silicon Valley  
2022/11/28 11:00:53 > 9294578d-a30c-416c-833d-42f294c94ed0  Season 1  2655dfa4-65cf-463e-af3f-7293ba6fd198  Silicon Valley  
2022/11/28 11:00:53 > ad02f2c0-483a-4d47-9a9e-574944dc58df  Season 4  2655dfa4-65cf-463e-af3f-7293ba6fd198  Silicon Valley  
2022/11/28 11:00:53 > 78e9d258-61d9-43f1-b068-8fc51a4626f6  Season 2  67c67307-a6af-4dd7-909d-cca3952a2373  IT Crowd  
2022/11/28 11:00:53 > 8288b3bb-ad47-4a11-933c-c9277ae122c2  Season 4  67c67307-a6af-4dd7-909d-cca3952a2373  IT Crowd  
2022/11/28 11:00:53 > 97fe1d67-9669-4ec9-9f85-77267f6ee927  Season 1  67c67307-a6af-4dd7-909d-cca3952a2373  IT Crowd  
2022/11/28 11:00:53 > c0f5455f-b417-49c9-be75-1631ed535ccc  Season 3  67c67307-a6af-4dd7-909d-cca3952a2373  IT Crowd  
2022/11/28 11:00:53 LEFT JOIN
2022/11/28 11:00:53 got 9 records
2022/11/28 11:00:53 > 16e5372f-f467-489e-8cc1-b16867befa6e  Season 5  2655dfa4-65cf-463e-af3f-7293ba6fd198  Silicon Valley  
2022/11/28 11:00:53 > 294701e8-9d72-4d87-8443-ca7f12985dc6  Season 3  2655dfa4-65cf-463e-af3f-7293ba6fd198  Silicon Valley  
2022/11/28 11:00:53 > 65f1547c-aac6-4ce9-af02-ca9cd809c893  Season 2  2655dfa4-65cf-463e-af3f-7293ba6fd198  Silicon Valley  
2022/11/28 11:00:53 > 9294578d-a30c-416c-833d-42f294c94ed0  Season 1  2655dfa4-65cf-463e-af3f-7293ba6fd198  Silicon Valley  
2022/11/28 11:00:53 > ad02f2c0-483a-4d47-9a9e-574944dc58df  Season 4  2655dfa4-65cf-463e-af3f-7293ba6fd198  Silicon Valley  
2022/11/28 11:00:53 > 78e9d258-61d9-43f1-b068-8fc51a4626f6  Season 2  67c67307-a6af-4dd7-909d-cca3952a2373  IT Crowd  
2022/11/28 11:00:53 > 8288b3bb-ad47-4a11-933c-c9277ae122c2  Season 4  67c67307-a6af-4dd7-909d-cca3952a2373  IT Crowd  
2022/11/28 11:00:53 > 97fe1d67-9669-4ec9-9f85-77267f6ee927  Season 1  67c67307-a6af-4dd7-909d-cca3952a2373  IT Crowd  
2022/11/28 11:00:53 > c0f5455f-b417-49c9-be75-1631ed535ccc  Season 3  67c67307-a6af-4dd7-909d-cca3952a2373  IT Crowd  
2022/11/28 11:00:53 RIGHT JOIN
2022/11/28 11:00:53 got 9 records
2022/11/28 11:00:53 > 16e5372f-f467-489e-8cc1-b16867befa6e  Season 5  2655dfa4-65cf-463e-af3f-7293ba6fd198  Silicon Valley  
2022/11/28 11:00:53 > 294701e8-9d72-4d87-8443-ca7f12985dc6  Season 3  2655dfa4-65cf-463e-af3f-7293ba6fd198  Silicon Valley  
2022/11/28 11:00:53 > 65f1547c-aac6-4ce9-af02-ca9cd809c893  Season 2  2655dfa4-65cf-463e-af3f-7293ba6fd198  Silicon Valley  
2022/11/28 11:00:53 > 9294578d-a30c-416c-833d-42f294c94ed0  Season 1  2655dfa4-65cf-463e-af3f-7293ba6fd198  Silicon Valley  
2022/11/28 11:00:53 > ad02f2c0-483a-4d47-9a9e-574944dc58df  Season 4  2655dfa4-65cf-463e-af3f-7293ba6fd198  Silicon Valley  
2022/11/28 11:00:53 > 78e9d258-61d9-43f1-b068-8fc51a4626f6  Season 2  67c67307-a6af-4dd7-909d-cca3952a2373  IT Crowd  
2022/11/28 11:00:53 > 8288b3bb-ad47-4a11-933c-c9277ae122c2  Season 4  67c67307-a6af-4dd7-909d-cca3952a2373  IT Crowd  
2022/11/28 11:00:53 > 97fe1d67-9669-4ec9-9f85-77267f6ee927  Season 1  67c67307-a6af-4dd7-909d-cca3952a2373  IT Crowd  
2022/11/28 11:00:53 > c0f5455f-b417-49c9-be75-1631ed535ccc  Season 3  67c67307-a6af-4dd7-909d-cca3952a2373  IT Crowd  
2022/11/28 11:00:53 FULL JOIN
2022/11/28 11:00:53 got 9 records
2022/11/28 11:00:53 > 16e5372f-f467-489e-8cc1-b16867befa6e  Season 5  2655dfa4-65cf-463e-af3f-7293ba6fd198  Silicon Valley  
2022/11/28 11:00:53 > 294701e8-9d72-4d87-8443-ca7f12985dc6  Season 3  2655dfa4-65cf-463e-af3f-7293ba6fd198  Silicon Valley  
2022/11/28 11:00:53 > 65f1547c-aac6-4ce9-af02-ca9cd809c893  Season 2  2655dfa4-65cf-463e-af3f-7293ba6fd198  Silicon Valley  
2022/11/28 11:00:53 > 9294578d-a30c-416c-833d-42f294c94ed0  Season 1  2655dfa4-65cf-463e-af3f-7293ba6fd198  Silicon Valley  
2022/11/28 11:00:53 > ad02f2c0-483a-4d47-9a9e-574944dc58df  Season 4  2655dfa4-65cf-463e-af3f-7293ba6fd198  Silicon Valley  
2022/11/28 11:00:53 > 78e9d258-61d9-43f1-b068-8fc51a4626f6  Season 2  67c67307-a6af-4dd7-909d-cca3952a2373  IT Crowd  
2022/11/28 11:00:53 > 8288b3bb-ad47-4a11-933c-c9277ae122c2  Season 4  67c67307-a6af-4dd7-909d-cca3952a2373  IT Crowd  
2022/11/28 11:00:53 > 97fe1d67-9669-4ec9-9f85-77267f6ee927  Season 1  67c67307-a6af-4dd7-909d-cca3952a2373  IT Crowd  
2022/11/28 11:00:53 > c0f5455f-b417-49c9-be75-1631ed535ccc  Season 3  67c67307-a6af-4dd7-909d-cca3952a2373  IT Crowd  
2022/11/28 11:00:53 LEFT SEMI JOIN
2022/11/28 11:00:53 got 9 records
2022/11/28 11:00:53 > 16e5372f-f467-489e-8cc1-b16867befa6e  Season 5  
2022/11/28 11:00:53 > 294701e8-9d72-4d87-8443-ca7f12985dc6  Season 3  
2022/11/28 11:00:53 > 65f1547c-aac6-4ce9-af02-ca9cd809c893  Season 2  
2022/11/28 11:00:53 > 78e9d258-61d9-43f1-b068-8fc51a4626f6  Season 2  
2022/11/28 11:00:53 > 8288b3bb-ad47-4a11-933c-c9277ae122c2  Season 4  
2022/11/28 11:00:53 > 9294578d-a30c-416c-833d-42f294c94ed0  Season 1  
2022/11/28 11:00:53 > 97fe1d67-9669-4ec9-9f85-77267f6ee927  Season 1  
2022/11/28 11:00:53 > ad02f2c0-483a-4d47-9a9e-574944dc58df  Season 4  
2022/11/28 11:00:53 > c0f5455f-b417-49c9-be75-1631ed535ccc  Season 3  
2022/11/28 11:00:53 Update table
2022/11/28 11:00:53 ok: after episodes table is updated
2022/11/28 11:00:53 > test 999 2016-06-26
2022/11/28 11:00:53 > test 999 2016-05-22
2022/11/28 11:00:53 > test 999 2015-05-03
2022/11/28 11:00:53 > test 999 2014-04-13
2022/11/28 11:00:53 > test 999 2017-05-21
2022/11/28 11:00:53 > test 999 2017-06-11
2022/11/28 11:00:53 > test 999 2017-06-04
2022/11/28 11:00:53 Delete Records
2022/11/28 11:00:53 > before delete: 9
2022/11/28 11:00:53 ok: deleted all records with first_aired from 2007-06-01 to 2008-06-01 in all seasons
2022/11/28 11:00:53 > after delete: 7
2022/11/28 11:00:53 Replace by fetch data
2022/11/28 11:00:53 > number of rows before replace: 70
2022/11/28 11:00:54 > number of rows after replace: 140
```