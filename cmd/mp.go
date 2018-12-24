package main

import (
	"encoding/json"
	"html"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/t-margheim/bcp-mp/pkg/opening"
)

var fakeResponse = html.UnescapeString(`<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
  <head>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0" />
    <title>oremus Bible Browser : Mark 1.1-11</title>
    <link rel="stylesheet" href="/bible.css" type="text/css" media="all" />
    <link rel="stylesheet" href="/bscreen.css" type="text/css" media="screen" />
    <link rel="stylesheet" href="/bprint.css" type="text/css" media="print" />
    <link rel="stylesheet" href="/obb/obb.css" type="text/css" media="all" />
    <link rev="made" href="mailto:biblemail&#64;oremus.org" />
    <script type="text/javascript" src="/bible.js?9"></script>
  </head>

  <body onLoad="FormLoad();">

    <script language="JavaScript">
      var ol_sticky=1;var ol_cap="footnote";
      ol_fixx=610;var ol_offsety=-25;ol_fgcolor="#FFEEEE";
      ol_bgcolor="#880000";ol_textcolor="black";ol_textsize="-1";ol_closecolor="#FFEEEE";
    </script>

    <div id="overDiv" style="position:absolute; visibility:hidden; z-index:1000;"></div>
    <script language="JavaScript" src="overlib_mini.js"><!-- overLIB (c) Erik Bosrup --></script>

    <script language="JavaScript">
      function setVis()
      {
        changeElementVisibility('sect');

      }
    </script>
<!--
<h1 align="center" id="h1screen" class="obb" style="font-size: 65px; color: #F00">oremus Bible Browser</h1>
-->
<h1 align="center" id="h1screen">
<span class="obb" style="font-size: 65px; color: #F00">oremus Bible Browser</span>
</h1>
<div class="quicklink">
<form method="post" action="/?version=NRSV&amp;passage=Mark%201.1-11" enctype="multipart/form-data">
<input type="submit" name="make_quicklink" value="make Quicklink" onmouseover="return overlib('Create a saved query and get an oBB Quicklink to it', 0, FIXX,-1, CAPTION,'', FGCOLOR,'#FFFCFC', BGCOLOR,'#CC3300', TEXTCOLOR,'#000000', OFFSETX,50, OFFSETY,-25, TEXTSIZE,-1, TIMEOUT,2500);" onmouseout="return nd();" />
<input type="hidden" name="passage" value="Mark 1.1-11"  />
<input type="hidden" name="version" value="NRSV"  />
</form>
</div>
<hr class="quicklink" /><div class="visbuttons"><nobr><label><input type="checkbox" name="vnum" value="no" checked="checked" onclick="changeElementVisibility('vnum')" id="vnum" /></label><label for="vnum">Omit&nbsp;verse&nbsp;numbers;</label></nobr>
<nobr><label><input type="checkbox" name="fnote" value="no" checked="checked" onclick="changeElementVisibility('fnote')" id="fnote" /></label><label for="fnote">Omit&nbsp;footnotes</label></nobr>
<br />
<nobr><label><input type="checkbox" name="headings" value="yes" checked="checked" onclick="changeElementVisibility('sect')" id="headings" /></label><label for="headings">Show&nbsp;section&nbsp;headings;</label></nobr>
<nobr><label><input type="checkbox" name="show_ref" value="no" onclick="changeElementVisibility('passageref')" id="ref" /></label><label for="ref">Omit&nbsp;passage&nbsp;reference</label></nobr>
<br />
<nobr><label><input type="checkbox" name="show_adj" value="no" onclick="changeElementVisibility('adj')" id="adj" /></label><label for="adj">Omit&nbsp;adjacent&nbsp;passage&nbsp;references</label></nobr>
<form onSubmit="removeHidden();return false;"><input type="submit" value="Remove hidden text" /></form><hr />
</div><!-- class="visbuttons" -->
<div class="bible">

<h2 class="passageref">Mark 1.1-11</h2>

<div class="bibletext">
<p>

<p><span class="cc">1</span>The beginning of the good news of Jesus Christ, the Son of God.
<p><sup class="ww">2</sup>As it is written in the prophet Isaiah, <blockquote>&#147;See, I am sending my messenger ahead of you,<br><spacer size=10>who will prepare your way;<br>
<sup class="ww">3</sup>the voice of one crying out in the wilderness:<br><spacer size=10>&#145;Prepare the way of the Lord,<br><spacer size=10>make his paths straight,&#146;&#148;</blockquote>
<sup class="ww">4</sup>John the baptizer appeared in the wilderness, proclaiming a baptism of repentance for the forgiveness of sins.
<sup class="ww">5</sup>And people from the whole Judean countryside and all the people of Jerusalem were going out to him, and were baptized by him in the river Jordan, confessing their sins.
<sup class="ww">6</sup>Now John was clothed with camel&#146;s hair, with a leather belt around his waist, and he ate locusts and wild honey.
<sup class="ww">7</sup>He proclaimed, &#147;The one who is more powerful than I is coming after me; I am not worthy to stoop down and untie the thong of his sandals.
<sup class="ww">8</sup>I have baptized you with water; but he will baptize you with the Holy Spirit.&#148;
<p><sup class="ww">9</sup>In those days Jesus came from Nazareth of Galilee and was baptized by John in the Jordan.
<sup class="ww">10</sup>And just as he was coming up out of the water, he saw the heavens torn apart and the Spirit descending like a dove on him.
<sup class="ww">11</sup>And a voice came from heaven, &#147;You are my Son, the Beloved; with you I am well pleased.&#148;
</p>

</div><!-- class="bibletext" -->

<div class="adj">
<table border="0" width="100%"><tr><td valign="top" align="right">&lt;&lt;</td><td valign="top" align="left"><form method="post" action="/?version=NRSV&amp;passage=Mark%201.1-11" enctype="multipart/form-data">
<input type="hidden" name="passage" value="Matthew 28" />
<input type="submit" name="show passage_button" value="Matthew 28" />
<input type="hidden" name="vnum" value="yes"  />
<input type="hidden" name="fnote" value="yes"  />
<input type="hidden" name="headings" value="no"  />
<input type="hidden" name="adj" value=""  />
<input type="hidden" name="version" value="NRSV" />
</form>
</td><td valign="top" align="right"><form method="post" action="/?version=NRSV&amp;passage=Mark%201.1-11" enctype="multipart/form-data">
<input type="hidden" name="passage" value="Mark 1.12-45" />
<input type="submit" name="show passage_button" value="Mark 1.12-45" />
<input type="hidden" name="vnum" value="yes"  />
<input type="hidden" name="fnote" value="yes"  />
<input type="hidden" name="headings" value="no"  />
<input type="hidden" name="adj" value=""  />
<input type="hidden" name="version" value="NRSV" />
</form>
</td><td valign="top" align="left">&gt;&gt;</td></tr></table>
</div><!-- class="adj" -->
</div><!-- class="bible" -->

<div class="copyright">
<hr />
<p>
<cite>New Revised Standard Version Bible</cite>, copyright &copy; 1989 National Council of the Churches of Christ in the United States of America. Used by permission. All rights reserved worldwide. <a href="http://nrsvbibles.org">http://nrsvbibles.org</a>
</p>
</div>
<div class="quicklink">
<form method="post" action="/?version=NRSV&amp;passage=Mark%201.1-11" enctype="multipart/form-data">
<input type="submit" name="make_quicklink" value="make Quicklink" onmouseover="return overlib('Create a saved query and get an oBB Quicklink to it', 0, FIXX,-1, CAPTION,'', FGCOLOR,'#FFFCFC', BGCOLOR,'#CC3300', TEXTCOLOR,'#000000', OFFSETX,50, OFFSETY,-25, TEXTSIZE,-1, TIMEOUT,2500);" onmouseout="return nd();" />
<input type="hidden" name="passage" value="Mark 1.1-11"  />
<input type="hidden" name="version" value="NRSV"  />
</form>
</div>
<hr class="quicklink" />
<div class="another">
<p>Enter another bible reference: </p><form method="post" action="/?version=NRSV&amp;passage=Mark%201.1-11" enctype="multipart/form-data">
<input type="text" name="passage" value="" size="22" maxlength="1024" /><input type="submit" name="show passage_button" value="show passage" />
<input type="hidden" name="vnum" value="yes"  />
<input type="hidden" name="fnote" value="yes"  />
<input type="hidden" name="headings" value="no"  />
<input type="hidden" name="adj" value=""  />
<input type="hidden" name="version" value="NRSV" />
</form>

</div> <!-- class="another" -->
    <div align="left" class="credits">
      <hr />
      <div class="screencredits">
	<p>
	  <a href="/">
<span class="obb" style="font-size: 68px; color: #F00; line-height: 0.65em;">obb</span><br />
<span class="obb" style="font-size: 18px; color: #F00">bible browser</span>
</a><br />
<!--
<img src="/obb1.gif" alt="oremus Bible Browser" border="0" width="81" height="70" /></a><br />
-->
	    <a href="mailto:biblemail&#64;oremus.org">biblemail&#64;oremus.org</a><br />
	  v&nbsp;2.2.8<br />
	  14 July 2018
	</p>
      </div>
      <div class="printcredits">
	From the oremus Bible Browser http://bible.oremus.org v2.2.8 14 July 2018.
      </div>
    </div>

  </body>

</html>
`)

func main() {
	// contents, err := ioutil.ReadFile("./gopher.json")
	// if err != nil {
	// 	log.Fatal("failed to read file:", err)
	// }
	// var story map[string]segment
	// err = json.Unmarshal(contents, &story)
	// if err != nil {
	// 	log.Fatal("failed to parse json:", err)
	// }

	// httpResponse, err := http.Get("http://bible.oremus.org/?version=NRSV&passage=Mark%201.1-11")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // rawResp, _ := ioutil.ReadAll(httpResponse.Body)
	// var response resp
	// decoder := xml.NewDecoder(httpResponse.Body)
	// // err = xml.Unmarshal(rawResp, &response)
	// decoder.Strict = false
	// err = decoder.Decode(&response)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Printf("%+v\n", response)
	app := prayerApp{}

	log.Fatal(http.ListenAndServe(":7777", &app))
}

type prayerApp struct {
}

func (a *prayerApp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req, _ := http.NewRequest(http.MethodGet, "https://api.esv.org/v3/passage/text?include-verse-numbers=false&q=Isa%2B42%3A1-12%3BEph%2B6%3A10-20%3BJohn%2B3%3A16-21&include-footnotes=false&include-headings=false", nil)
	req.Header.Add("Authorization", "Token a9a234f364de585a1a6273b00ffe4be9c1b9ab47")
	httpResponse, _ := http.DefaultClient.Do(req)
	responseBody, _ := ioutil.ReadAll(httpResponse.Body)

	var response resp
	err := json.Unmarshal(responseBody, &response)
	if err != nil {
		log.Println(err)
	}

	log.Printf("%+v\n", response)

	date := time.Now().Add(-7 * time.Hour)

	selectedDate := r.URL.Query().Get("date")
	if selectedDate != "" {
		newDate, err := time.Parse("2006-01-02", selectedDate)
		if err != nil {
			log.Println(err.Error())
		}
		date = newDate
	}

	open, _ := opening.Get(date)
	elements := content{
		Date:    date.Format("2006-01-02"),
		Opening: open,
		Invitatory: invitatory{
			Name: "Venite",
			Content: `Come, let us sing to the Lord; *<br>
			let us shout for joy to the Rock of our salvation. 
		Let us come before his presence with thanksgiving * 
			and raise a loud shout to him with psalms.
		
		For the Lord is a great God, * 
			and a great King above all gods. 
		In his hand are the caverns of the earth, * 
			and the heights of the hills are his also. 
		The sea is his, for he made it, * 
			and his hands have molded the dry land.
		
		Come, let us bow down, and bend the knee, * 
			and kneel before the Lord our Maker. 
		For he is our God, 
		and we are the people of his pasture and the sheep of his hand. *
			Oh, that today you would hearken to his voice!`,
		},
		GospelText: response.Passages[2],
		Lesson1:    response.Passages[0],
		Lesson2:    response.Passages[1],
	}
	template := template.Must(template.ParseFiles("./mp.html"))

	template.Execute(w, elements)
	return
}

type content struct {
	Date       string
	Opening    opening.Opening
	Invitatory invitatory
	GospelText string
	Lesson1    string
	Lesson2    string
}

type invitatory struct {
	Name    string
	Content template.HTML
}

type resp struct {
	Passages []string
}
