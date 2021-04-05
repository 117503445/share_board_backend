package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("H4sIAAAAAAAC/wrwZmYRYeBg4GBwlWkNYkACIgycDMXlienpqUX6UFovqzg/LzSElYHRuDUr4U3Ew2wZV5H/9/YnRfydcGhj9fUtHN9PNgTkdk+79oBd9+2dd6fvXbk99+WZnH6N7wsnZOneEpnF1Rmhuy1ARKrZaYvqk/k779r+5755c+/TvPJnVTNWBOjNN8uv+zzb/P3fe9Uza+qYFFt0Lh7fx38nS/n1lLtbpHZObcm+krVYZrGOHMd0YTVOXuM3phePyqwr03h27ypfmX3rhd1tUWu3nzOZ7j/frvnSXF6R+/ek9/1QvWvnsc/eZcbEppiLnkrVEwQzKy6cOvpdbaMer2rZ0yn2WYpHGh9p9M+9sZRB8H3q61jNyNuFv/7ui3XK9YwU5WNUy2rWScz/PfNsWWbC1BsHW/JtHzlpHHhz++b38Is3u8UvbLp7xfgOi17Xgm1vpu2ds5ir5eiGN10vfSy+Hq5Z4a5+/9xK8Z0TDGtiokM5zkT077A4vbo49sKFVQfXmbzaI/Tb2GL3hdcXf3a4iLDMWJ++XV427+h9W6f3q85KabAfzNBwyFi/SmH5LN/HSpdW/Gj+PKvZa0PBboEHM1u4vqu1b+15y3u82VRoUk+sRXL52+Kmtuld8a8cJf/zbNEufb0ldcHUKc2ShUoLV1XFLeUoUPZ7sstwtfGiqneL+Tfnhj/b0+THrsfrKiKzP+D/fKaAjfO7053+v7y2+cm5fy/u1e35+fbF6fmP/+4tSU+u+GGxc+tuDfbJW4P4pradW2Fwonbhdqf8k/nnH+nvu8bZk9jHPOHAWfbz+mtDGf8YfZh9pPX0Ef7wm6o7DL1zHn+8OZNVqWuX896vkRp85kqHDm6+5fGJ459pK/uV5nlr976YmLrO++kixiu662e0+/X/8jKZunLHC+HQlLX+S27rP25u/+85PyLxidO05dd3uvrc+3jCdMMLPrXVG61XrP+onH1o4svl0+zit35TvD4zSOltDHt4WPzUnmC2GCFnLZ221hfB72YKmz4TnzT5wnr31AUn68Neubb+CQ4LfiO1i+Fq2j0lL9X2uDvHZRSmVxmvC9R8un/bNvfdu1/2b582bdbbJ+f+/njHz3sl67j/N7f99rkL/BelFUtGvRUq+7Rk4Rfv5Z4aB6/7isfesNU5O3fRq+x/9Y1PHq4/aX5+elv9myPl8urf5x2bbX/0fHlbfZ3o/XvycjYnpl+ezj/JXXr9l60cH2LWsR2af/tyyue6X09aP55s61ui/tX0XcTj67+sl52ban+Sl+lwz5z2Vw46z196qeyeMvn8lflyoQF2/rsb9p6MET83V+KPLAPD//8B3uwctnN3H21mZmAIFWFggOVmBozczI7IzfAMDNKNrCbAm5FJhBlRGiCbDCoNYGBbI4jEWzYgjMLuFAgQYPjveI6ZAYvDWNlA8kwMTAydDAwMPCwgHiAAAP//LgLG/qsEAAA="); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
