package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("H4sIAAAAAAAC/wrwZmYRYeBg4GBYpWkVxIAERBg4GYrLE9PTU4v0obReVnF+XmgIKwPjhyX8CW8iHmbLuIr8v7c/KeLvhEeFu69v4fh+smHCLbG1uc72GRuLd+tXB+eXLzPebLvoTrvTqq1XllwI8ThyLYnbdaGhhkuw87+Zt/d+nvy/6c2dt5tuzt0pt0xryunlZ2ff//PsbPX+73Fq8+WdORS9Sza+Z94+64RVYPmaWye8DqVmmV7kWKwjxzFdWI2TN/hN6MVjMnaZgWf2rn5hsmw6B9/FdhbebWGBNdnL93+/nFBs//Pdg/X27471f/tVnZ0xUa9YhqGKO7Zt2oJ3yxWOH1/JrDJR1mMl95mMh3+qXXu8nTtjK08YbxV45PA21+Se3dMz820qV23wZlS653n9dtCmZbnlu+9W8a31PnWqgMMgUDqFr/q4f27pdLawvreiNa8FxJ0Y0+6L3Y936n5+fuHPXYvTdzDG9By+vcJaMqtng8S03EIL9bKvBvKrLkrmPr/SHeMhabdqjRZL2hrnmDlpN/apZWSEJoRnhcZ1/uWck/ui5EWJixIX05nwS7f5+c0e5csvupy1me2Uo4EUl8Ct1VonOI8b6in1zrO3WG2pwc69nZdxnseKgvyiCL+b9zbYRqgpet5snbsz2l45oyLxj7epxk1//91BW+W2fdkyaZOhcplKSK3/ahHGvCOdMYZppU4h52P2TcnbbfRUtMSqIFNj0YOr3F//azKI/l5ZwfVLPsfGO+893/Wyx1/z2+L/Sbx/x7Mt7Dd7/+TZtzcXGUR6Pw3IetHD1XwnvUr6te7Pp5F/jsbOPSDhx76dJc+uoipTl/OBeON9Bb0YlbodR/QOfPVMOn0sPVuoQ+OV2NcwtYUWMf0c7OezlE4q/7guIhfAb1X61rVfpzhWK513QVblSZst1lfCF2n6Ho2cr6GsXZu04afSvfoP8790srqdW1pZNTfa5WOf+7rGyHKDyN67jrWp/WualNvnR/2qCb3VHpPo1bmllDcmsi7UyE+ywJNn5QYJsUXeO9P9Qo7pq6gvjO/18p/8f8uV46Kf/CN9N65+yRQys3KzxlOZqsL9jxh2Xpld0Buk+X/nRss79+Rq6ovzzmWqbX9/tFzuasBiSfubkv9+bWmz65NJl5M6K1P2Z8fCL8XLPTUOXvcVj71hq3N27qI3r1/8ZzrZPH/y+vnnv9bf7npv43dt3rHX9kfOb/+1v170/q237+/L6/LH3fGXW7fvyvXChuVFpgL/rsxTe/2dzf9bj79ZiYrtyvVrklPXfy33nZv9RcWcg8e4SOwUs+H85dwTc4L97ILrq6Y4Lq68xfZGtNi27F3E+/cMDAz//wd4s3PMW7F/9SxmBoZ4EQYGWCZnwMjk7IhMDs/XIN3IagK8GZlEmBGFBLLJoEICBrY1gki8RQbCKOxOgQABhv+OT5kZsDiMlQ0kz8TAxNDJwMCgzALiAQIAAP//N9Qk+cIEAAA="); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
