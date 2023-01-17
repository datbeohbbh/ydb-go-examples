```text
go run ./basic/xorm -ydb=grpcs://ydb:2135/local -prefix=xorm -showSQL=false
2023/01/17 21:38:32 ok: connected to database
2023/01/17 21:38:32 ok: xorm/test/episodes table is created
2023/01/17 21:38:32 ok: xorm/series table is created
2023/01/17 21:38:33 ok: xorm/seasons table is created
2023/01/17 21:38:33 ok: xorm/episodes table is created
2023/01/17 21:38:33 ok: fill tables with data
2023/01/17 21:38:33 Replace by fetch data
2023/01/17 21:38:33 > number of rows of `xorm/test/episodes`: 70
2023/01/17 21:38:33 > number of rows after replace: 70
2023/01/17 21:38:33 Select Default
2023/01/17 21:38:33 > select of all known series:
2023/01/17 21:38:33 > [538c44de-b50a-4025-a237-7cfdb87b7fce] Silicon Valley (2014-04-06)
2023/01/17 21:38:33 > [7b5f494c-dccd-4549-bba7-0a008452bacb] IT Crowd (2006-02-03)
2023/01/17 21:38:33 Select Scan
2023/01/17 21:38:33 > scan select of episodes of `Season 1` of `IT Crowd` between 2006-01-01 and 2006-12-31:
2023/01/17 21:38:33 > [56afaa75-bbe3-436a-ab80-8a57e13651d3] Calamity Jen (2006-02-03)
2023/01/17 21:38:33 > [f15db472-f1f7-4ff7-a792-9d45f259bd08] Yesterday's Jam (2006-02-03)
2023/01/17 21:38:33 > [5cc350b0-4f05-4a1a-a933-073af6d91c15] Fifty-Fifty (2006-02-10)
2023/01/17 21:38:33 > [f5363b3c-12d7-4b62-84b9-fc1d419f1491] The Red Door (2006-02-17)
2023/01/17 21:38:33 > [1a14b73b-3c8f-4fc2-8d3f-1b9d1ee1dd4c] The Haunting of Bill Crouse (2006-02-24)
2023/01/17 21:38:33 > [877b47e5-32a2-43b5-a795-21c1878acbfb] Aunt Irma Visits (2006-03-03)
2023/01/17 21:38:33 INNER JOIN
2023/01/17 21:38:33 got 9 records
2023/01/17 21:38:33 > 181128e8-10fe-4d1c-a2d9-1066d77f5f72  Season 2  538c44de-b50a-4025-a237-7cfdb87b7fce  Silicon Valley  
2023/01/17 21:38:33 > 59044ef4-e072-4c37-a5ba-3345dc89d367  Season 1  538c44de-b50a-4025-a237-7cfdb87b7fce  Silicon Valley  
2023/01/17 21:38:33 > 6aaf17ec-cbef-48f6-b94c-f51b5d05c5ce  Season 5  538c44de-b50a-4025-a237-7cfdb87b7fce  Silicon Valley  
2023/01/17 21:38:33 > c5f65a7e-dd0b-46ad-879c-c531ab4149b5  Season 4  538c44de-b50a-4025-a237-7cfdb87b7fce  Silicon Valley  
2023/01/17 21:38:33 > d5f2fb5e-92f6-415d-b1b9-2babc2503841  Season 3  538c44de-b50a-4025-a237-7cfdb87b7fce  Silicon Valley  
2023/01/17 21:38:33 > 59222bb7-3c78-4596-ba2c-68693d92f365  Season 4  7b5f494c-dccd-4549-bba7-0a008452bacb  IT Crowd  
2023/01/17 21:38:33 > 64b5ce29-fce8-4462-8b3b-ba508e3206d9  Season 3  7b5f494c-dccd-4549-bba7-0a008452bacb  IT Crowd  
2023/01/17 21:38:33 > 7ff75d5b-2833-4f60-8ec3-01cb136b591b  Season 1  7b5f494c-dccd-4549-bba7-0a008452bacb  IT Crowd  
2023/01/17 21:38:33 > 85b6ef72-6fa8-4222-95f3-a44c40a524fc  Season 2  7b5f494c-dccd-4549-bba7-0a008452bacb  IT Crowd  
2023/01/17 21:38:33 LEFT JOIN
2023/01/17 21:38:33 got 9 records
2023/01/17 21:38:33 > 181128e8-10fe-4d1c-a2d9-1066d77f5f72  Season 2  538c44de-b50a-4025-a237-7cfdb87b7fce  Silicon Valley  
2023/01/17 21:38:33 > 59044ef4-e072-4c37-a5ba-3345dc89d367  Season 1  538c44de-b50a-4025-a237-7cfdb87b7fce  Silicon Valley  
2023/01/17 21:38:33 > 6aaf17ec-cbef-48f6-b94c-f51b5d05c5ce  Season 5  538c44de-b50a-4025-a237-7cfdb87b7fce  Silicon Valley  
2023/01/17 21:38:33 > c5f65a7e-dd0b-46ad-879c-c531ab4149b5  Season 4  538c44de-b50a-4025-a237-7cfdb87b7fce  Silicon Valley  
2023/01/17 21:38:33 > d5f2fb5e-92f6-415d-b1b9-2babc2503841  Season 3  538c44de-b50a-4025-a237-7cfdb87b7fce  Silicon Valley  
2023/01/17 21:38:33 > 59222bb7-3c78-4596-ba2c-68693d92f365  Season 4  7b5f494c-dccd-4549-bba7-0a008452bacb  IT Crowd  
2023/01/17 21:38:33 > 64b5ce29-fce8-4462-8b3b-ba508e3206d9  Season 3  7b5f494c-dccd-4549-bba7-0a008452bacb  IT Crowd  
2023/01/17 21:38:33 > 7ff75d5b-2833-4f60-8ec3-01cb136b591b  Season 1  7b5f494c-dccd-4549-bba7-0a008452bacb  IT Crowd  
2023/01/17 21:38:33 > 85b6ef72-6fa8-4222-95f3-a44c40a524fc  Season 2  7b5f494c-dccd-4549-bba7-0a008452bacb  IT Crowd  
2023/01/17 21:38:34 RIGHT JOIN
2023/01/17 21:38:34 got 9 records
2023/01/17 21:38:34 > 181128e8-10fe-4d1c-a2d9-1066d77f5f72  Season 2  538c44de-b50a-4025-a237-7cfdb87b7fce  Silicon Valley  
2023/01/17 21:38:34 > 59044ef4-e072-4c37-a5ba-3345dc89d367  Season 1  538c44de-b50a-4025-a237-7cfdb87b7fce  Silicon Valley  
2023/01/17 21:38:34 > 6aaf17ec-cbef-48f6-b94c-f51b5d05c5ce  Season 5  538c44de-b50a-4025-a237-7cfdb87b7fce  Silicon Valley  
2023/01/17 21:38:34 > c5f65a7e-dd0b-46ad-879c-c531ab4149b5  Season 4  538c44de-b50a-4025-a237-7cfdb87b7fce  Silicon Valley  
2023/01/17 21:38:34 > d5f2fb5e-92f6-415d-b1b9-2babc2503841  Season 3  538c44de-b50a-4025-a237-7cfdb87b7fce  Silicon Valley  
2023/01/17 21:38:34 > 59222bb7-3c78-4596-ba2c-68693d92f365  Season 4  7b5f494c-dccd-4549-bba7-0a008452bacb  IT Crowd  
2023/01/17 21:38:34 > 64b5ce29-fce8-4462-8b3b-ba508e3206d9  Season 3  7b5f494c-dccd-4549-bba7-0a008452bacb  IT Crowd  
2023/01/17 21:38:34 > 7ff75d5b-2833-4f60-8ec3-01cb136b591b  Season 1  7b5f494c-dccd-4549-bba7-0a008452bacb  IT Crowd  
2023/01/17 21:38:34 > 85b6ef72-6fa8-4222-95f3-a44c40a524fc  Season 2  7b5f494c-dccd-4549-bba7-0a008452bacb  IT Crowd  
2023/01/17 21:38:34 FULL JOIN
2023/01/17 21:38:34 got 9 records
2023/01/17 21:38:34 > 181128e8-10fe-4d1c-a2d9-1066d77f5f72  Season 2  538c44de-b50a-4025-a237-7cfdb87b7fce  Silicon Valley  
2023/01/17 21:38:34 > 59044ef4-e072-4c37-a5ba-3345dc89d367  Season 1  538c44de-b50a-4025-a237-7cfdb87b7fce  Silicon Valley  
2023/01/17 21:38:34 > 6aaf17ec-cbef-48f6-b94c-f51b5d05c5ce  Season 5  538c44de-b50a-4025-a237-7cfdb87b7fce  Silicon Valley  
2023/01/17 21:38:34 > c5f65a7e-dd0b-46ad-879c-c531ab4149b5  Season 4  538c44de-b50a-4025-a237-7cfdb87b7fce  Silicon Valley  
2023/01/17 21:38:34 > d5f2fb5e-92f6-415d-b1b9-2babc2503841  Season 3  538c44de-b50a-4025-a237-7cfdb87b7fce  Silicon Valley  
2023/01/17 21:38:34 > 59222bb7-3c78-4596-ba2c-68693d92f365  Season 4  7b5f494c-dccd-4549-bba7-0a008452bacb  IT Crowd  
2023/01/17 21:38:34 > 64b5ce29-fce8-4462-8b3b-ba508e3206d9  Season 3  7b5f494c-dccd-4549-bba7-0a008452bacb  IT Crowd  
2023/01/17 21:38:34 > 7ff75d5b-2833-4f60-8ec3-01cb136b591b  Season 1  7b5f494c-dccd-4549-bba7-0a008452bacb  IT Crowd  
2023/01/17 21:38:34 > 85b6ef72-6fa8-4222-95f3-a44c40a524fc  Season 2  7b5f494c-dccd-4549-bba7-0a008452bacb  IT Crowd  
2023/01/17 21:38:34 LEFT SEMI JOIN
2023/01/17 21:38:34 got 9 records
2023/01/17 21:38:34 > 181128e8-10fe-4d1c-a2d9-1066d77f5f72  Season 2  
2023/01/17 21:38:34 > 59044ef4-e072-4c37-a5ba-3345dc89d367  Season 1  
2023/01/17 21:38:34 > 59222bb7-3c78-4596-ba2c-68693d92f365  Season 4  
2023/01/17 21:38:34 > 64b5ce29-fce8-4462-8b3b-ba508e3206d9  Season 3  
2023/01/17 21:38:34 > 6aaf17ec-cbef-48f6-b94c-f51b5d05c5ce  Season 5  
2023/01/17 21:38:34 > 7ff75d5b-2833-4f60-8ec3-01cb136b591b  Season 1  
2023/01/17 21:38:34 > 85b6ef72-6fa8-4222-95f3-a44c40a524fc  Season 2  
2023/01/17 21:38:34 > c5f65a7e-dd0b-46ad-879c-c531ab4149b5  Season 4  
2023/01/17 21:38:34 > d5f2fb5e-92f6-415d-b1b9-2babc2503841  Season 3  
2023/01/17 21:38:34 Update table
2023/01/17 21:38:34 ok: after episodes table is updated
2023/01/17 21:38:34 > test 999 2015-05-03
2023/01/17 21:38:34 > test 999 2014-04-13
2023/01/17 21:38:34 > test 999 2017-06-11
2023/01/17 21:38:34 > test 999 2017-06-04
2023/01/17 21:38:34 > test 999 2017-05-21
2023/01/17 21:38:34 > test 999 2016-05-22
2023/01/17 21:38:34 > test 999 2016-06-26
2023/01/17 21:38:34 Delete Records
2023/01/17 21:38:34 > before delete: 9
2023/01/17 21:38:34 ok: deleted all records with first_aired from 2007-06-01 to 2008-06-01 in all seasons
2023/01/17 21:38:34 > after delete: 7
2023/01/17 21:38:34 Close scheme engine
2023/01/17 21:38:34 Close data engine
2023/01/17 21:38:34 Close scan engine
```