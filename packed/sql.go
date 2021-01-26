package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("H4sIAAAAAAAC/wrwZmYRYeBg4GDwUbMKYkACYgycDMlFqYklqXrFhTl6+gh2aAgrA+PE+fwJb6af9ntsI3L8+9RffOkzvRRYa+8nua3onjBHrPuKY9/ixqmPzdo6l/lOeGxw5nb/nG1TvDmXinP2nlhVWvPv+a/NtyUsA/JbDK88e3Y758//zO2X39646nhkAs/KHQeyH85fXC/oLH/n2dePhjLbbc6y6R0PZp168euRK1EHfVc6Tg/e2lh6YJH3g3uRHQ5sDeEqzwyXXLOwfVtW0yZ3/Ojvn1oNl98emO0/K12z5/PtPQYTLszg0WRzUOYz4XT+ICTH6VTbprzkz4rQjEVTTbSELy2Z4LnYxy74XP3enQENzPqcz1e9Nfoau41/po/vVbkULstgTZ/Tj66tYGf+t/O4zXxlZmvW7pUM9lzO0d+Olj+sK49VLDY8ef97wckTJ5KvFzRef3Aj7YtCJ6fsIcn0jrXLWtgqUqV/KOQsPzp5ytzlP+2D35RVNZR4aDjJewROOXizW89Z3i9O5Vb91Kz62E06vxTZBKUrryn9CD4RaXezs3NXUNQblmOLD2Vbha2/58gya+OW3bcjMnNDlS+aZKp3Pj+8iedlhurenx8OZDXk1XUejzrD/ShH84W7ybMzylPWB8fOm9X9IVxPgm+C4tfCykm6nh0aLBJK7gkSxW41edMnnT5QvWlLbci3y6dXTHOtlbta89x9UdTxm88bs5hfb6+M6r9ULVIba/L2YoiX/J9Nnu9TQv/qlZZF6H+Kk1oZvGN61afgrOdHSu/eK+szW/+gTu8S2//690p7J29fdetAje2Xa79Mv5x9Jf5l94kbQl2L5vW+/ZiZ/cXxd6Fm6bU/SyxUzs00+bRA7ltznvHj+TPuTqvvU19+99x9BgaG//8DvNk5OCzve4UxMTDk8jEwwJIoA8NKTdQkyoWSRMHJ8v0S/gSQAcjKArwZmUSYEakc2XBQKoeBbY0gkkCaRxiG3T0QIMDw33ExEwN217GygZQwMTAx9DEwMDxhAvEAAQAA//9Ooi+OiAMAAA=="); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
