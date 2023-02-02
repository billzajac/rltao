package rltao

import (
	"testing"
)

type passagetestcase struct {
	in    int
	title string
	body  string
}

func TestGetPassage(t *testing.T) {
	// Beware of tabs in the body!
	cases := []passagetestcase{
		{1, "The Beginning of Power", `The Tao that can be expressed<br>
	Is not the Tao of the Absolute.<br>
	The name that can be named<br>
	Is not the name of the Absolute.<br>
	<br>
	The nameless originated Heaven and Earth.<br>
	The named is the Mother of All Things.<br>
	<br>
	Thus, without expectation<br>
	One will always perceive the subtlety;<br>
	And, with expectation<br>
	One will always perceive the boundary.<br>
	<br>
	The source of these two is identical,<br>
	Yet their names are different.<br>
	Together they are called profound,<br>
	Profound and mysterious,<br>
	The gateway to the Collective Subtlety.`},
		{81, "The Evolved Way", `Sincere words are not embellished;<br>
	Embellished words are not sincere.<br>
	Those who are good are not defensive;<br>
	Those who are defensive are not good.<br>
	Those who know are not erudite;<br>
	Those who are erudite do not know.<br>
	<br>
	Evolved Individuals do not accumulate.<br>
	The more they do for others, the more they gain;<br>
	The more they give to others, the more they possess.<br>
	<br>
	The Tao of Nature<br>
	Is to serve without spoiling.<br>
	The Tao of Evolved Individuals<br>
	Is to act wihtout contending.`},
		{15, "The Power in Subtle Force", `Those skillful in the ancient Tao<br>
	Are subtly ingenious and profoundly intuitive.<br>
	They are so deep they cannot be recognized.<br>
	Since, indeed, they cannot be recognized,<br>
	Their force can be contained.<br>
	<br>
	So careful!<br>
	As if wading a stream in winter.<br>
	So hesitant!<br>
	As if respecting all sides in the community.<br>
	So reserved!<br>
	As if acting as a guest.<br>
	So yielding!<br>
	As if ice about to melt.<br>
	So candid!<br>
	As if acting wiht simplicity.<br>
	So open!<br>
	As if acting as a valley.<br>
	So integrated!<br>
	As if acting as muddy water.<br>
	<br>
	Who can harmonize with muddy water,<br>
	And gradually arrive at clarity?<br>
	Who can move with stability,<br>
	And gradually bring endurance to life?<br>
	<br>
	Those who maintain the Tao<br>
	Do not desire to become full.<br>
	Indeed, since they are not full,<br>
	They can be used up and also renewed.`},
	}
	for _, c := range cases {
		got := GetPassage(c.in)
		if got.Title != c.title {
			t.Errorf("GenerateTetragram(%d) == %q, want title: %q", c.in, got.Title, c.title)
		}
		if got.Body != c.body {
			t.Errorf("GenerateTetragram(%d)\n GOT:%q\nWANT:%q", c.in, got.Body, c.body)
		}
	}
}
