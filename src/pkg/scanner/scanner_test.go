package scanner_test

import (
	. "scanner";
	"testing";
	"strings";
	"fmt";
)

func TestScanner1(t *testing.T) {
	sc := NewScanner(strings.NewReader("1 2 3"));
	
	for sc.HasNextInt64() {
		fmt.Println(sc.NextInt64());
	}

	sc = NewScanner(strings.NewReader("1 2 -3"));

	for sc.HasNextUint64() {
		fmt.Println(sc.NextUint64());
	}
}

func TestScanner2(t *testing.T) {
	input := `    125  6   00 9139081 -1309714037 1037104 -0183  
	 
	  
	   091 apa        kabar 0      `;

	{
		s := strings.NewReader(input);
		sc := NewScanner(s);
		
		for sc.HasNextInt() {
			fmt.Printf("int:#%d#\n", sc.NextInt());
		}
	}
	{
		sc := NewScannerString(input);
		
		fmt.Printf("uint:#%d#\n", sc.NextUint());
		fmt.Printf("uint:#%d#\n", sc.NextUint());
		
		for sc.HasNextLine() {
			fmt.Println("line:#" + sc.NextLine() + "#");
		}
	}
	{
		sc := NewScannerString(input);
		
		for sc.HasNext() {
			fmt.Println("token:#" + sc.Next() + "#");
		}
	}
}

func testHasNext(s string) {
	sc := NewScannerString(s);
	
	fmt.Println(s, "has int =", sc.HasNextInt());
	fmt.Println(s, "has int64 =", sc.HasNextInt64());
	fmt.Println(s, "has uint =", sc.HasNextUint());
	fmt.Println(s, "has uint64 =", sc.HasNextUint64());
	fmt.Println(s, "has line =", sc.HasNextLine());
}

func TestScanner3(t *testing.T) {
	testHasNext("123");
	testHasNext("-123");
	testHasNext("123456789123456789");
	testHasNext("-123456789123456789");
}

