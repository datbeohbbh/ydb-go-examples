// Code generated by ydbgen; DO NOT EDIT.

package main

import (
	"strconv"

	"github.com/ydb-platform/ydb-go-sdk/v3"
	"github.com/ydb-platform/ydb-go-sdk/v3/table"
)

var (
	_ = strconv.Itoa
	_ = ydb.StringValue
	_ = table.NewQueryParameters
)

func (i *Item) Scan(res *table.Result) (err error) {
	res.SeekItem("host_uid")
	i.HostUID = res.OUint64()

	res.SeekItem("url_uid")
	i.URLUID = res.OUint64()

	res.SeekItem("url")
	i.URL = res.OUTF8()

	res.SeekItem("page")
	i.Page = res.OUTF8()

	return res.Err()
}

func (i *Item) StructValue() ydb.Value {
	var v0 ydb.Value
	{
		var v1 ydb.Value
		{
			vp0 := ydb.OptionalValue(ydb.Uint64Value(i.HostUID))
			v1 = vp0
		}
		var v2 ydb.Value
		{
			vp0 := ydb.OptionalValue(ydb.Uint64Value(i.URLUID))
			v2 = vp0
		}
		var v3 ydb.Value
		{
			vp0 := ydb.OptionalValue(ydb.UTF8Value(i.URL))
			v3 = vp0
		}
		var v4 ydb.Value
		{
			vp0 := ydb.OptionalValue(ydb.UTF8Value(i.Page))
			v4 = vp0
		}
		v0 = ydb.StructValue(
			ydb.StructFieldValue("host_uid", v1),
			ydb.StructFieldValue("url_uid", v2),
			ydb.StructFieldValue("url", v3),
			ydb.StructFieldValue("page", v4),
		)
	}
	return v0
}

func (is *ItemList) Scan(res *table.Result) (err error) {
	for res.NextRow() {
		var x0 Item
		res.SeekItem("host_uid")
		x0.HostUID = res.OUint64()

		res.SeekItem("url_uid")
		x0.URLUID = res.OUint64()

		res.SeekItem("url")
		x0.URL = res.OUTF8()

		res.SeekItem("page")
		x0.Page = res.OUTF8()

		if res.Err() == nil {
			*is = append(*is, x0)
		}
	}
	return res.Err()
}

func (is ItemList) ListValue() ydb.Value {
	var list0 ydb.Value
	vs0 := make([]ydb.Value, len(is))
	for i0, item0 := range is {
		var v0 ydb.Value
		{
			var v1 ydb.Value
			{
				var v2 ydb.Value
				{
					vp0 := ydb.OptionalValue(ydb.Uint64Value(item0.HostUID))
					v2 = vp0
				}
				var v3 ydb.Value
				{
					vp0 := ydb.OptionalValue(ydb.Uint64Value(item0.URLUID))
					v3 = vp0
				}
				var v4 ydb.Value
				{
					vp0 := ydb.OptionalValue(ydb.UTF8Value(item0.URL))
					v4 = vp0
				}
				var v5 ydb.Value
				{
					vp0 := ydb.OptionalValue(ydb.UTF8Value(item0.Page))
					v5 = vp0
				}
				v1 = ydb.StructValue(
					ydb.StructFieldValue("host_uid", v2),
					ydb.StructFieldValue("url_uid", v3),
					ydb.StructFieldValue("url", v4),
					ydb.StructFieldValue("page", v5),
				)
			}
			v0 = v1
		}
		vs0[i0] = v0
	}
	if len(vs0) == 0 {
		var t1 ydb.Type
		{
			var t2 ydb.Type
			{
				fs0 := make([]ydb.StructOption, 4)
				var t3 ydb.Type
				{
					tp0 := ydb.TypeUint64
					t3 = ydb.Optional(tp0)
				}
				fs0[0] = ydb.StructField("host_uid", t3)
				var t4 ydb.Type
				{
					tp0 := ydb.TypeUint64
					t4 = ydb.Optional(tp0)
				}
				fs0[1] = ydb.StructField("url_uid", t4)
				var t5 ydb.Type
				{
					tp0 := ydb.TypeUTF8
					t5 = ydb.Optional(tp0)
				}
				fs0[2] = ydb.StructField("url", t5)
				var t6 ydb.Type
				{
					tp0 := ydb.TypeUTF8
					t6 = ydb.Optional(tp0)
				}
				fs0[3] = ydb.StructField("page", t6)
				t2 = ydb.Struct(fs0...)
			}
			t1 = t2
		}
		t0 := ydb.List(t1)
		list0 = ydb.ZeroValue(t0)
	} else {
		list0 = ydb.ListValue(vs0...)
	}
	return list0
}
