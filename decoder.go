package base64

import "unsafe"

//go:nosplit
func (e *Encoding) decode(dst []byte, src []byte) int {
	inlen := uintptr(len(src))
	if inlen == 0 || (e.pad && (inlen&3) != 0) {
		return 0
	}
	ip := (*sliceHeader)(unsafe.Pointer(&src)).data
	ipstart := ip
	op := (*sliceHeader)(unsafe.Pointer(&dst)).data
	opstart := op
	var cu uint32
	if inlen >= 8+4 {
		ux := ctou32(ip)
		vx := ctou32(ip + 4)
		for ip < (ipstart+inlen)-(128+4) {
			{
				_u := ux
				ux = ctou32(ip + 8 + 0*8)
				_u = (e.lutXd0[byte(_u)] | e.lutXd1[byte(_u>>8)] | e.lutXd2[byte(_u>>16)] | e.lutXd3[_u>>24])
				cu |= _u
				stou32(op+0*6, _u)
				_v := vx
				vx = ctou32(ip + 8 + 0*8 + 4)
				_v = (e.lutXd0[byte(_v)] | e.lutXd1[byte(_v>>8)] | e.lutXd2[byte(_v>>16)] | e.lutXd3[_v>>24])
				stou32(op+0*6+3, _v)
			}
			{
				_u := ux
				ux = ctou32(ip + 8 + 1*8)
				stou32(op+1*6, (e.lutXd0[byte(_u)] | e.lutXd1[byte(_u>>8)] | e.lutXd2[byte(_u>>16)] | e.lutXd3[_u>>24]))
				_v := vx
				vx = ctou32(ip + 8 + 1*8 + 4)
				stou32(op+1*6+3, (e.lutXd0[byte(_v)] | e.lutXd1[byte(_v>>8)] | e.lutXd2[byte(_v>>16)] | e.lutXd3[_v>>24]))
			}
			{
				_u := ux
				ux = ctou32(ip + 8 + 2*8)
				stou32(op+2*6, (e.lutXd0[byte(_u)] | e.lutXd1[byte(_u>>8)] | e.lutXd2[byte(_u>>16)] | e.lutXd3[_u>>24]))
				_v := vx
				vx = ctou32(ip + 8 + 2*8 + 4)
				stou32(op+2*6+3, (e.lutXd0[byte(_v)] | e.lutXd1[byte(_v>>8)] | e.lutXd2[byte(_v>>16)] | e.lutXd3[_v>>24]))
			}
			{
				_u := ux
				ux = ctou32(ip + 8 + 3*8)
				stou32(op+3*6, (e.lutXd0[byte(_u)] | e.lutXd1[byte(_u>>8)] | e.lutXd2[byte(_u>>16)] | e.lutXd3[_u>>24]))
				_v := vx
				vx = ctou32(ip + 8 + 3*8 + 4)
				stou32(op+3*6+3, (e.lutXd0[byte(_v)] | e.lutXd1[byte(_v>>8)] | e.lutXd2[byte(_v>>16)] | e.lutXd3[_v>>24]))
			}
			{
				_u := ux
				ux = ctou32(ip + 8 + 4*8)
				stou32(op+4*6, (e.lutXd0[byte(_u)] | e.lutXd1[byte(_u>>8)] | e.lutXd2[byte(_u>>16)] | e.lutXd3[_u>>24]))
				_v := vx
				vx = ctou32(ip + 8 + 4*8 + 4)
				stou32(op+4*6+3, (e.lutXd0[byte(_v)] | e.lutXd1[byte(_v>>8)] | e.lutXd2[byte(_v>>16)] | e.lutXd3[_v>>24]))
			}
			{
				_u := ux
				ux = ctou32(ip + 8 + 5*8)
				stou32(op+5*6, (e.lutXd0[byte(_u)] | e.lutXd1[byte(_u>>8)] | e.lutXd2[byte(_u>>16)] | e.lutXd3[_u>>24]))
				_v := vx
				vx = ctou32(ip + 8 + 5*8 + 4)
				stou32(op+5*6+3, (e.lutXd0[byte(_v)] | e.lutXd1[byte(_v>>8)] | e.lutXd2[byte(_v>>16)] | e.lutXd3[_v>>24]))
			}
			{
				_u := ux
				ux = ctou32(ip + 8 + 6*8)
				stou32(op+6*6, (e.lutXd0[byte(_u)] | e.lutXd1[byte(_u>>8)] | e.lutXd2[byte(_u>>16)] | e.lutXd3[_u>>24]))
				_v := vx
				vx = ctou32(ip + 8 + 6*8 + 4)
				stou32(op+6*6+3, (e.lutXd0[byte(_v)] | e.lutXd1[byte(_v>>8)] | e.lutXd2[byte(_v>>16)] | e.lutXd3[_v>>24]))
			}
			{
				_u := ux
				ux = ctou32(ip + 8 + 7*8)
				stou32(op+7*6, (e.lutXd0[byte(_u)] | e.lutXd1[byte(_u>>8)] | e.lutXd2[byte(_u>>16)] | e.lutXd3[_u>>24]))
				_v := vx
				vx = ctou32(ip + 8 + 7*8 + 4)
				stou32(op+7*6+3, (e.lutXd0[byte(_v)] | e.lutXd1[byte(_v>>8)] | e.lutXd2[byte(_v>>16)] | e.lutXd3[_v>>24]))
			}

			{
				_u := ux
				ux = ctou32(ip + 8 + 8*8)
				stou32(op+8*6, (e.lutXd0[byte(_u)] | e.lutXd1[byte(_u>>8)] | e.lutXd2[byte(_u>>16)] | e.lutXd3[_u>>24]))
				_v := vx
				vx = ctou32(ip + 8 + 8*8 + 4)
				stou32(op+8*6+3, (e.lutXd0[byte(_v)] | e.lutXd1[byte(_v>>8)] | e.lutXd2[byte(_v>>16)] | e.lutXd3[_v>>24]))
			}
			{
				_u := ux
				ux = ctou32(ip + 8 + 9*8)
				stou32(op+9*6, (e.lutXd0[byte(_u)] | e.lutXd1[byte(_u>>8)] | e.lutXd2[byte(_u>>16)] | e.lutXd3[_u>>24]))
				_v := vx
				vx = ctou32(ip + 8 + 9*8 + 4)
				stou32(op+9*6+3, (e.lutXd0[byte(_v)] | e.lutXd1[byte(_v>>8)] | e.lutXd2[byte(_v>>16)] | e.lutXd3[_v>>24]))
			}
			{
				_u := ux
				ux = ctou32(ip + 8 + 10*8)
				stou32(op+10*6, (e.lutXd0[byte(_u)] | e.lutXd1[byte(_u>>8)] | e.lutXd2[byte(_u>>16)] | e.lutXd3[_u>>24]))
				_v := vx
				vx = ctou32(ip + 8 + 10*8 + 4)
				stou32(op+10*6+3, (e.lutXd0[byte(_v)] | e.lutXd1[byte(_v>>8)] | e.lutXd2[byte(_v>>16)] | e.lutXd3[_v>>24]))
			}
			{
				_u := ux
				ux = ctou32(ip + 8 + 11*8)
				stou32(op+11*6, (e.lutXd0[byte(_u)] | e.lutXd1[byte(_u>>8)] | e.lutXd2[byte(_u>>16)] | e.lutXd3[_u>>24]))
				_v := vx
				vx = ctou32(ip + 8 + 11*8 + 4)
				stou32(op+11*6+3, (e.lutXd0[byte(_v)] | e.lutXd1[byte(_v>>8)] | e.lutXd2[byte(_v>>16)] | e.lutXd3[_v>>24]))
			}
			{
				_u := ux
				ux = ctou32(ip + 8 + 12*8)
				stou32(op+12*6, (e.lutXd0[byte(_u)] | e.lutXd1[byte(_u>>8)] | e.lutXd2[byte(_u>>16)] | e.lutXd3[_u>>24]))
				_v := vx
				vx = ctou32(ip + 8 + 12*8 + 4)
				stou32(op+12*6+3, (e.lutXd0[byte(_v)] | e.lutXd1[byte(_v>>8)] | e.lutXd2[byte(_v>>16)] | e.lutXd3[_v>>24]))
			}
			{
				_u := ux
				ux = ctou32(ip + 8 + 13*8)
				stou32(op+13*6, (e.lutXd0[byte(_u)] | e.lutXd1[byte(_u>>8)] | e.lutXd2[byte(_u>>16)] | e.lutXd3[_u>>24]))
				_v := vx
				vx = ctou32(ip + 8 + 13*8 + 4)
				stou32(op+13*6+3, (e.lutXd0[byte(_v)] | e.lutXd1[byte(_v>>8)] | e.lutXd2[byte(_v>>16)] | e.lutXd3[_v>>24]))
			}
			{
				_u := ux
				ux = ctou32(ip + 8 + 14*8)
				stou32(op+14*6, (e.lutXd0[byte(_u)] | e.lutXd1[byte(_u>>8)] | e.lutXd2[byte(_u>>16)] | e.lutXd3[_u>>24]))
				_v := vx
				vx = ctou32(ip + 8 + 14*8 + 4)
				stou32(op+14*6+3, (e.lutXd0[byte(_v)] | e.lutXd1[byte(_v>>8)] | e.lutXd2[byte(_v>>16)] | e.lutXd3[_v>>24]))
			}
			{
				_u := ux
				ux = ctou32(ip + 8 + 15*8)
				stou32(op+15*6, (e.lutXd0[byte(_u)] | e.lutXd1[byte(_u>>8)] | e.lutXd2[byte(_u>>16)] | e.lutXd3[_u>>24]))
				_v := vx
				vx = ctou32(ip + 8 + 15*8 + 4)
				stou32(op+15*6+3, (e.lutXd0[byte(_v)] | e.lutXd1[byte(_v>>8)] | e.lutXd2[byte(_v>>16)] | e.lutXd3[_v>>24]))
			}
			ip += 128
			op += (128 / 4) * 3
		}
		for ip < (ipstart+inlen)-(16+4) {
			{
				_u := ux
				ux = ctou32(ip + 8 + 0*8)
				_u = (e.lutXd0[byte(_u)] | e.lutXd1[byte(_u>>8)] | e.lutXd2[byte(_u>>16)] | e.lutXd3[_u>>24])
				cu |= _u
				stou32(op+0*6, _u)
				_v := vx
				vx = ctou32(ip + 8 + 0*8 + 4)
				stou32(op+0*6+3, (e.lutXd0[byte(_v)] | e.lutXd1[byte(_v>>8)] | e.lutXd2[byte(_v>>16)] | e.lutXd3[_v>>24]))
			}
			{
				_u := ux
				ux = ctou32(ip + 8 + 1*8)
				stou32(op+1*6, (e.lutXd0[byte(_u)] | e.lutXd1[byte(_u>>8)] | e.lutXd2[byte(_u>>16)] | e.lutXd3[_u>>24]))
				_v := vx
				vx = ctou32(ip + 8 + 1*8 + 4)
				stou32(op+1*6+3, (e.lutXd0[byte(_v)] | e.lutXd1[byte(_v>>8)] | e.lutXd2[byte(_v>>16)] | e.lutXd3[_v>>24]))
			}
			ip += 16
			op += (16 / 4) * 3
		}
		if ip < (ipstart+inlen)-(8+4) {
			stou32(op+0*6, (e.lutXd0[byte(ux)] | e.lutXd1[byte(ux>>8)] | e.lutXd2[byte(ux>>16)] | e.lutXd3[ux>>24]))
			stou32(op+0*6+3, (e.lutXd0[byte(vx)] | e.lutXd1[byte(vx>>8)] | e.lutXd2[byte(vx>>16)] | e.lutXd3[vx>>24]))
			ip += 8
			op += (8 / 4) * 3
		}
	}
	for ip < (ipstart+inlen)-4 {
		u := ctou32(ip)
		u = (e.lutXd0[byte(u)] | e.lutXd1[byte(u>>8)] | e.lutXd2[byte(u>>16)] | e.lutXd3[u>>24])
		stou32(op, u)
		cu |= u
		ip += 4
		op += 3
	}
	var u uint32
	l := (ipstart + inlen) - ip
	if l == 4 {
		if *(*byte)(unsafe.Pointer(ip + 3)) == '=' {
			l = 3
			if *(*byte)(unsafe.Pointer(ip + 2)) == '=' {
				l = 2
				if *(*byte)(unsafe.Pointer(ip + 1)) == '=' {
					l = 1
				}
			}
		}
	}
	up := (*[3]byte)(unsafe.Pointer(&u))
	switch l {
	case 4:
		u = ctou32(ip)
		u = (e.lutXd0[byte(u)] | e.lutXd1[byte(u>>8)] | e.lutXd2[byte(u>>16)] | e.lutXd3[u>>24])
		*(*byte)(unsafe.Pointer(op)) = up[0]
		op++
		*(*byte)(unsafe.Pointer(op)) = up[1]
		op++
		*(*byte)(unsafe.Pointer(op)) = up[2]
		op++
		cu |= u
		break
	case 3:
		u = e.lutXd0[*(*byte)(unsafe.Pointer(ip + 0))] | e.lutXd1[*(*byte)(unsafe.Pointer(ip + 1))] | e.lutXd2[*(*byte)(unsafe.Pointer(ip + 2))]
		*(*byte)(unsafe.Pointer(op)) = up[0]
		op++
		*(*byte)(unsafe.Pointer(op)) = up[1]
		op++
		cu |= u
		break
	case 2:
		u = e.lutXd0[*(*byte)(unsafe.Pointer(ip + 0))] | e.lutXd1[*(*byte)(unsafe.Pointer(ip + 1))]
		*(*byte)(unsafe.Pointer(op)) = up[0]
		op++
		cu |= u
		break
	case 1:
		u = e.lutXd0[*(*byte)(unsafe.Pointer(ip + 0))]
		*(*byte)(unsafe.Pointer(op)) = up[0]
		op++
		cu |= u
		break
	}
	if cu == 0xffffffff {
		return 0
	}
	return int(op - opstart)
}
