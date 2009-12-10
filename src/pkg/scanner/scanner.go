package main

import (
	"fmt";
	"strings";
	"io";
	"bufio";
	"bytes";
	"unicode";
	"os";
	"container/list";
	"strconv";
)

type Scanner struct {
	in *bufio.Reader;
	list *list.List;
}

type sd struct {
	s string;
	delim int;
	err os.Error;
}

func NewScanner(r io.Reader) *Scanner {
	s := Scanner{in: bufio.NewReader(r)};
	s.list = list.New();
	return &s;
}

func (this *Scanner) nextToken() (s string, delim int, err os.Error) {
	buf := bytes.NewBufferString("");
	
	for {
		if c, _, e := this.in.ReadRune(); e == nil {
			if unicode.IsSpace(c) {
				s = buf.String();
				delim = c;
				return; // (token, delim, nil)
			} else {
				buf.WriteString(string(c));
			}
		} else {
			if e == os.EOF {
				if buf.Len() > 0 {
					s = buf.String();
					return; // (token, 0, nil)
				}
			}
			err = e;
			return; // ("", 0, os.EOF)
		}
	}
	
	return;
}

func (this *Scanner) nextBuffedToken() sd {
	if this.list.Len() == 0 {
		s, delim, err := this.nextToken();
		next := sd{s: s, delim: delim, err: err};
		this.list.PushBack(next);
	}
	
	return this.list.Front().Value.(sd);
}

func (this *Scanner) popBuff() {
	if this.list.Len() == 0 {
		panicln("should not pop list with len 0");
	}
	this.list.Remove(this.list.Front());
}

func (this *Scanner) NextInt() int {
	for {
		next := this.nextBuffedToken();

		if next.err != nil {
			panicln("Error encountered. Call Has* funcs before calling this");
		} else {
			defer this.popBuff();
			
			if len(next.s) > 0 {
				// yeah!
				if v, e := strconv.Atoi(next.s); e == nil {
					return v;
				} else {
					panicln("String data was " + next.s + ". Cannot convert to int");
				}
			}
		}
	}
	panicln("should not reach here");
	return 0;
}

func (this *Scanner) HasNextInt() bool {
	for {
		next := this.nextBuffedToken();
		
		if next.err != nil {
			return false;
		}
		
		if len(next.s) > 0 {
			// we have the data, check if it's an int
			
			_, e := strconv.Atoi(next.s);
			if e == nil {
				return true;
			} else {
				return false;
			}
		}
		
		// last was double-delimiter. so we go back to loop after removing first element.
		this.popBuff();
	}
	panicln("should not reach here");
	return false;
}

func main() {
	input := `    125  6   00 9139081 1309714037 1037104 0183
	 
	  
	   091 apa        kabar 0      `;

	s := strings.NewReader(input);
	sc := NewScanner(s);
	
	for sc.HasNextInt() {
		fmt.Println(sc.NextInt());
	}
	
	println("beres");
}
