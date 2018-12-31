package prayers

import "html/template"

type Prayer struct {
	Name    string
	Content template.HTML
}

var DailyPrayers = [][]Prayer{
	[]Prayer{ // 1-14-18-27
		Prayer{
			Name: "For Joy in God's Creation",
			Content: `O heavenly Father, <em>who hast</em> filled the world with beauty:
			Open our eyes to behold <em>thy</em> gracious hand in all <em>thy</em> works;
			that, rejoicing in <em>thy</em> whole creation, we may learn to serve
			<em>thee</em> with gladness; for the sake of him through whom all
			things were made, <em>thy</em> Son Jesus Christ our Lord. <em>Amen.</em>`,
		},
		Prayer{
			Name: "For the Unity of the Church",
			Content: `O God the Father of our Lord Jesus Christ, our only Savior,
			the Prince of Peace: Give us grace seriously to lay to heart the
			great dangers we are in by our unhappy divisions; take away
			all hatred and prejudice, and whatever else may hinder us
			from godly union and concord; that, as there is but one Body
			and one Spirit, one hope of our calling, one Lord, one Faith,
			one Baptism, one God and Father of us all, so we may be all
			of one heart and of one soul, united in one holy bond of truth
			and peace, of faith and charity, and may with one mind and
			one mouth glorify <em>thee</em>; through Jesus Christ our Lord.
			<em>Amen.</em>`,
		},
		Prayer{
			Name: "For our Country",
			Content: `Almighty God, who hast given us this good land for our
			heritage: We humbly beseech thee that we may always prove
			ourselves a people mindful of thy favor and glad to do thy will.
			Bless our land with honorable industry, sound learning, and
			pure manners. Save us from violence, discord, and confusion;
			from pride and arrogance, and from every evil way. Defend
			our liberties, and fashion into one united people the multitudes
			brought hither out of many kindreds and tongues. Endue
			with the spirit of wisdom those to whom in thy Name we entrust
			the authority of government, that there may be justice and
			peace at home, and that, through obedience to thy law, we
			may show forth thy praise among the nations of the earth.
			In the time of prosperity, fill our hearts with thankfulness,
			and in the day of trouble, suffer not our trust in thee to fail;
			all which we ask through Jesus Christ our Lord. <em>Amen.</em>`,
		},
		Prayer{
			Name: "For Social Justice",
			Content: `Grant, O God, that your holy and life-giving Spirit may so
			move every human heart [and especially the hearts of the
			people of this land], that barriers which divide us may
			crumble, suspicions disappear, and hatreds cease; that our
			divisions being healed, we may live in justice and peace;
			through Jesus Christ our Lord. <em>Amen.</em>`,
		},
	},
	[]Prayer{ // 2-16-19-28
		Prayer{
			Name: "For All Sorts and Conditions of Men",
			Content: `O God, the creator and preserver of all mankind, we humbly
			beseech thee for all sorts and conditions of men; that thou
			wouldest be pleased to make thy ways known unto them, thy
			saving health unto all nations. More especially we pray for
			thy holy Church universal; that it may be so guided and
			governed by thy good Spirit, that all who profess and call
			themselves Christians may be led into the way of truth, and
			hold the faith in unity of spirit, in the bond of peace, and in
			righteousness of life. Finally, we commend to thy fatherly
			goodness all those who are in any ways afflicted or distressed,
			in mind, body, or estate; [especially those for whom our prayers
			are desired]; that it may please thee to comfort and relieve
			them according to their several necessities, giving them patience
			under their sufferings, and a happy issue out of all their
			afflictions. And this we beg for Jesus Christ's sake. <em>Amen.</em>`,
		},
		Prayer{
			Name: "For Monastic Orders and Vocations",
			Content: `O Lord Jesus Christ, you became poor for our sake, that we
			might be made rich through your poverty: Guide and sanctify,
			we pray, those whom you call to follow you under the vows
			of poverty, chastity, and obedience, that by their prayer and
			service they may enrich your Church, and by their life and
			worship may glorify your Name; for you reign with the Father
			and the Holy Spirit, one God, now and for ever. <em>Amen.</em>`,
		},
		Prayer{
			Name: "For the President of the United States and all in Civil Authority",
			Content: ` Lord our Governor, whose glory is in all the world: We
			commend this nation to <em>thy</em> merciful care, that, being guided
			by <em>thy</em> Providence, we may dwell secure in <em>thy</em> peace. Grant
			to the President of the United States, the Governor of this
			State (<em>or</em> Commonwealth), and to all in authority, wisdom
			and strength to know and to do <em>thy</em> will. Fill them with the
			love of truth and righteousness, and make them ever mindful
			of their calling to serve this people in <em>thy</em> fear; through Jesus
			Christ our Lord, who <em>liveth</em> and <em>reigneth</em> with <em>thee</em> and the
			Holy Spirit, one God, world without end. <em>Amen.</em>`,
		},
		Prayer{
			Name: "In Times of Conflict",
			Content: `O God, you have bound us together in a common life. Help us,
			in the midst of our struggles for justice and truth, to confront
			one another without hatred or bitterness, and to work
			together with mutual forbearance and respect; through Jesus
			Christ our Lord. <em>Amen.</em>`,
		},
	},
	[]Prayer{ // 3-17-20-29
		Prayer{
			Name: "For the Human Family",
			Content: `O God, you made us in your own image and redeemed us
			through Jesus your Son: Look with compassion on the whole
			human family; take away the arrogance and hatred which
			infect our hearts; break down the walls that separate us;
			unite us in bonds of love; and work through our struggle and
			confusion to accomplish your purposes on earth; that, in
			your good time, all nations and races may serve you in
			harmony around your heavenly throne; through Jesus Christ
			our Lord. <em>Amen.</em>`,
		},
		Prayer{
			Name: "For Church Musicians and Artists",
			Content: `O God, whom saints and angels delight to worship in
			heaven: Be ever present with your servants who seek through
			art and music to perfect the praises offered by your people on
			earth; and grant to them even now glimpses of your beauty,
			and make them worthy at length to behold it unveiled for
			evermore; through Jesus Christ our Lord. <em>Amen.</em></p>`,
		},
		Prayer{
			Name: "For Congress or a State Legislature",
			Content: `O God, the fountain of wisdom, whose will is good and
			gracious, and whose law is truth: We beseech <em>thee</em> so to guide
			and bless our Senators and Representatives in Congress
			assembled (<em>or</em> in the Legislature of this State, <em>or</em> Common-
			wealth), that they may enact such laws as shall please <em>thee</em>,
			to the glory of <em>thy</em> Name and the welfare of this people;
			through Jesus Christ our Lord. <em>Amen.</em>`,
		},
		Prayer{
			Name: "For Agriculture",
			Content: `Almighty God, we thank you for making the earth fruitful, so
			that it might produce what is needed for life: Bless those who
			work in the fields; give us seasonable weather; and grant that
			we may all share the fruits of the earth, rejoicing in your
			goodness; through Jesus Christ our Lord. <em>Amen.</em>`,
		},
	},
	[]Prayer{ // 4-21-30-33
		Prayer{
			Name: "For Peace",
			Content: `Eternal God, in whose perfect kingdom no sword is drawn
			but the sword of righteousness, no strength known but the
			strength of love: So mightily spread abroad your Spirit, that
			all peoples may be gathered under the banner of the Prince of
			Peace, as children of one Father; to whom be dominion and
			glory, now and for ever. <em>Amen.</em>`,
		},
		Prayer{
			Name: "For Courts of Justice",
			Content: `Almighty God, <em>who sittest</em> in the throne judging right: We
			humbly beseech <em>thee</em> to bless the courts of justice and the
			magistrates in all this land; and give unto them the spirit of
			wisdom and understanding, that they may discern the truth,
			and impartially administer the law in the fear of <em>thee</em> alone;
			through him who shall come to be our Judge, <em>thy</em> Son our
			Savior Jesus Christ. <em>Amen.</em>`,
		},
		Prayer{
			Name: "For the Unemployed",
			Content: `Heavenly Father, we remember before you those who suffer
			want and anxiety from lack of work. Guide the people of this
			land so to use our public and private wealth that all may find
			suitable and fulfilling employment, and receive just payment
			for their labor; through Jesus Christ our Lord. <em>Amen.</em>`,
		},
		Prayer{
			Name: "For Cities",
			Content: `Heavenly Father, in your Word you have given us a vision of
			that holy City to which the nations of the world bring their
			glory: Behold and visit, we pray, the cities of the earth.
			Renew the ties of mutual regard which form our civic life.
			Send us honest and able leaders. Enable us to eliminate
			poverty, prejudice, and oppression, that peace may prevail
			with righteousness, and justice with order, and that men and
			women from different cultures and with differing talents may
			find with one another the fulfillment of their humanity;
			through Jesus Christ our Lord. <em>Amen.</em>`,
		},
	},
	[]Prayer{ // 5-56-25-31
		Prayer{
			Name: "For Peace Among the Nations",
			Content: `Almighty God our heavenly Father, guide the nations of the
			world into the way of justice and truth, and establish among
			them that peace which is the fruit of righteousness, that they
			may become the kingdom of our Lord and Savior Jesus Christ.
			<em>Amen.</em>`,
		},
		Prayer{
			Name: "For the Victims of Addiction",
			Content: `O blessed Lord, you ministered to all who came to you: Look
			with compassion upon all who through addiction have lost
			their health and freedom. Restore to them the assurance of
			your unfailing mercy; remove from them the fears that beset
			them; strengthen them in the work of their recovery; and to
			those who care for them, give patient understanding and
			persevering love. <em>Amen.</em></br>`,
		},
		Prayer{
			Name: "For those in the Armed Forces of our Country",
			Content: `Almighty God, we commend to your gracious care and
			keeping all the men and women of our armed forces at home
			and abroad. Defend them day by day with your heavenly
			grace; strengthen them in their trials and temptations; give
			them courage to face the perils which beset them; and grant
			them a sense of your abiding presence wherever they may be;
			through Jesus Christ our Lord. <em>Amen.</em>`,
		},
		Prayer{
			Name: "For Schools and Colleges",
			Content: `O Eternal God, bless all schools, colleges, and universities
			[and especially ___________], that they may be lively centers for
			sound learning, new discovery, and the pursuit of wisdom;
			and grant that those who teach and those who learn may find
			you to be the source of all truth; through Jesus Christ our
			Lord. <em>Amen.</em>`,
		},
	},
	[]Prayer{ // 6-23-26-32
		Prayer{
			Name: "For our Enemies",
			Content: `O God, the Father of all, whose Son commanded us to love
			our enemies: Lead them and us from prejudice to truth:
			deliver them and us from hatred, cruelty, and revenge; and in
			your good time enable us all to stand reconciled before you,
			through Jesus Christ our Lord. <em>Amen.</em>`,
		},
		Prayer{
			Name: "For Local Government",
			Content: `Almighty God our heavenly Father, send down upon those
			who hold office in this State (Commonwealth, City, County,
			Town, ____________) the spirit of wisdom, charity, and justice;
			that with steadfast purpose they may faithfully serve in their
			offices to promote the well-being of all people; through Jesus
			Christ our Lord. <em>Amen.</em>`,
		},
		Prayer{
			Name: "For those who suffer for the sake of Conscience",
			Content: `O God our Father, whose Son forgave his enemies while he
			was suffering shame and death: Strengthen those who suffer
			for the sake of conscience; when they are accused, save them
			from speaking in hate; when they are rejected, save them
			from bitterness; when they are imprisoned, save them from
			despair; and to us your servants, give grace to respect their
			witness and to discern the truth, that our society may be
			cleansed and strengthened. This we ask for the sake of Jesus
			Christ, our merciful and righteous Judge. <em>Amen.</em>`,
		},
		Prayer{
			Name: "For the Good Use of Leisure",
			Content: `O God, in the course of this busy life, give us times of
			refreshment and peace; and grant that we may so use our
			leisure to rebuild our bodies and renew our minds, that our
			spirits may be opened to the goodness of your creation;
			through Jesus Christ our Lord. <em>Amen.</em>`,
		},
	},
	[]Prayer{ // 7-36-38-40
		Prayer{
			Name: "For the Church",
			Content: `Gracious Father, we pray for thy holy Catholic Church. Fill it
			with all truth, in all truth with all peace. Where it is corrupt,
			purify it; where it is in error, direct it; where in any thing it is
			amiss, reform it. Where it is right, strengthen it; where it is in
			want, provide for it; where it is divided, reunite it; for the sake
			of Jesus Christ thy Son our Savior. <em>Amen.</em>`,
		},
		Prayer{
			Name: "For the Oppressed",
			Content: `Look with pity, O heavenly Father, upon the people in this
			land who live with injustice, terror, disease, and death as
			their constant companions. Have mercy upon us. Help us to
			eliminate our cruelty to these our neighbors. Strengthen those
			who spend their lives establishing equal protection of the law
			and equal opportunities for all. And grant that every one of
			us may enjoy a fair portion of the riches of this land; through
			Jesus Christ our Lord. <em>Amen.</em>`,
		},
		Prayer{
			Name: "For the Right Use of God's Gifts",
			Content: `Almighty God, whose loving hand <em>hath</em> given us all that we
			possess: Grant us grace that we may honor <em>thee</em> with our
			substance, and, remembering the account which we must one
			day give, may be faithful stewards of <em>thy</em> bounty, through
			Jesus Christ our Lord. <em>Amen.</em>`,
		},
		Prayer{
			Name: "For Knowledge of God's Creation",
			Content: `Almighty and everlasting God, you made the universe with
			all its marvelous order, its atoms, worlds, and galaxies, and
			the infinite complexity of living creatures: Grant that, as we
			probe the mysteries of your creation, we may come to know
			you more truly, and more surely fulfill our role in your
			eternal purpose; in the name of Jesus Christ our Lord. <em>Amen.</em>`,
		},
	},
	[]Prayer{ // 8-34-37-41
		Prayer{
			Name: "For the Mission of the Church",
			Content: `Everliving God, whose will it is that all should come to you
			through your Son Jesus Christ: Inspire our witness to him,
			that all may know the power of his forgiveness and the hope
			of his resurrection; who lives and reigns with you and the
			Holy Spirit, one God, now and for ever. <em>Amen.</em>`,
		},
		Prayer{
			Name: "For Towns and Rural Areas",
			Content: `Lord Christ, when you came among us, you proclaimed the
			kingdom of God in villages, towns, and lonely places: Grant
			that your presence and power may be known throughout this
			land. Have mercy upon all of us who live and work in rural
			areas [especially ___________]; and grant that all the people
			of our nation may give thanks to you for food and drink and
			all other bodily necessities of life, respect those who labor to
			produce them, and honor the land and the water from which
			these good things come. All this we ask in your holy Name.
			<em>Amen.</em></br>`,
		},
		Prayer{
			Name: "For Prisons and Correctional Institutions",
			Content: `Lord Jesus, for our sake you were condemned as a criminal:
			Visit our jails and prisons with your pity and judgment.
			Remember all prisoners, and bring the guilty to repentance
			and amendment of life according to your will, and give them
			hope for their future. When any are held unjustly, bring them
			release; forgive us, and teach us to improve our justice.
			Remember those who work in these institutions; keep them
			humane and compassionate; and save them from becoming
			brutal or callous. And since what we do for those in prison,
			O Lord, we do for you, constrain us to improve their lot. All
			this we ask for your mercy's sake. <em>Amen.</em></br>`,
		},
		Prayer{
			Name: "For the Conservation of Natural Resources",
			Content: `Almighty God, in giving us dominion over things on earth,
			you made us fellow workers in your creation: Give us wisdom
			and reverence so to use the resources of nature, that no one
			may suffer from our abuse of them, and that generations yet
			to come may continue to praise you for your bounty; through
			Jesus Christ our Lord. <em>Amen.</em></br>`,
		},
	},
	[]Prayer{ // 9-35-42-47
		Prayer{
			Name: "Clergy and People",
			Content: `Almighty and everlasting God, from whom cometh every
			good and perfect gift: Send down upon our bishops, and
			other clergy, and upon the congregations committed to their
			charge, the healthful Spirit of thy grace: and, that they may
			truly please thee, pour upon them the continual dew of thy
			blessing. Grant this, O Lord, for the honor of our Advocate
			and Mediator, Jesus Christ. <em>Amen.</em>`,
		},
		Prayer{
			Name: "For the Poor and the Neglected",
			Content: `Almighty and most merciful God, we remember before you
			all poor and neglected persons whom it would be easy for us
			to forget: the homeless and the destitute, the old and the sick,
			and all who have none to care for them. Help us to heal those
			who are broken in body or spirit, and to turn their sorrow
			into joy. Grant this, Father, for the love of your Son, who for
			our sake became poor, Jesus Christ our Lord. <em>Amen</em>.`,
		},
		Prayer{
			Name: "For the Harvest of Lands and Waters",
			Content: `O gracious Father, <em>who openest thine</em> hand and <em>fillest</em> all
			things living with plenteousness: Bless the lands and waters,
			and multiply the harvests of the world; let <em>thy</em> Spirit go
			forth, that it may renew the face of the earth; show <em>thy</em>
			loving-kindness, that our land may give her increase; and
			save us from selfish use of what <em>thou givest</em>, that men and
			women everywhere may give <em>thee</em> thanks; through Christ 
			our Lord. <em>Amen.</em>`,
		},
		Prayer{
			Name: "For Young Persons",
			Content: `God our Father, you see your children growing up in an
			unsteady and confusing world: Show them that your ways
			give more life than the ways of the world, and that following
			you is better than chasing after selfish goals. Help them to
			take failure, not as a measure of their worth, but as a chance
			for a new start. Give them strength to hold their faith in you,
			and to keep alive their joy in your creation; through Jesus
			Christ our Lord. <em>Amen.</em>`,
		},
	},
	[]Prayer{ // 10-39-44-45
		Prayer{
			Name: "For the Diocese",
			Content: `O God, by your grace you have called us in this Diocese to a
			goodly fellowship of faith. Bless our Bishops(s) <em>N.</em> [and <em>N.</em>],
			and other clergy, and all our people. Grant that your Word
			may be truly preached and truly heard, your Sacraments
			faithfully administered and faithfully received. By your
			Spirit, fashion our lives according to the example of your
			Son, and grant that we may show the power of your love to
			all among whom we live; through Jesus Christ our Lord.
			<em>Amen.</em>`,
		},
		Prayer{
			Name: "For those who Influence Public Opinion",
			Content: `Almighty God, you proclaim your truth in every age by many
			voices: Direct, in our time, we pray, those who speak where
			many listen and write what many read; that they may do their
			part in making the heart of this people wise, its mind sound, and
			its will righteous; to the honor of Jesus Christ our Lord. <em>Amen.</em>`,
		},
		Prayer{
			Name: "For the Future of the Human Race",
			Content: `O God our heavenly Father, you have blessed us and given us
			dominion over all the earth: Increase our reverence before
			the mystery of life; and give us new insight into your purposes
			for the human race, and new wisdom and determination in
			making provision for its future in accordance with your will;
			through Jesus Christ our Lord. <em>Amen.</em>`,
		},
		Prayer{
			Name: "For Families",
			Content: `Almighty God, our heavenly Father, who settest the solitary
			in families: We commend to thy continual care the homes in
			which thy people dwell. Put far from them, we beseech thee,
			every root of bitterness, the desire of vainglory, and the pride
			of life. Fill them with faith, virtue, knowledge, temperance,
			patience, godliness. Knit together in constant affection those
			who, in holy wedlock, have been made one flesh. Turn the
			hearts of the parents to the children, and the hearts of the
			children to the parents; and so enkindle fervent charity among
			us all, that we may evermore be kindly affectioned one
			to another; through Jesus Christ our Lord. <em>Amen.</em>`,
		},
	},
	[]Prayer{ // 11-46-48-49
		Prayer{
			Name: "For the Parish",
			Content: `Almighty and everliving God, ruler of all things in heaven
			and earth, hear our prayers for this parish family. Strengthen
			the faithful, arouse the careless, and restore the penitent.
			Grant us all things necessary for our common life, and bring
			us all to be of one heart and mind within your holy Church;
			through Jesus Christ our Lord. <em>Amen.</em>`,
		},
		Prayer{
			Name: "For the Care of Children",
			Content: `Almighty God, heavenly Father, you have blessed us with the
			joy and care of children: Give us calm strength and patient
			wisdom as we bring them up, that we may teach them to love
			whatever is just and true and good, following the example of
			our Savior Jesus Christ. <em>Amen.</em>`,
		},
		Prayer{
			Name: "For Those Who Live Alone",
			Content: `Almighty God, whose Son had nowhere to lay his head:
			Grant that those who live alone may not be lonely in their
			solitude, but that, following in his steps, they may find
			fulfillment in loving you and their neighbors; through Jesus
			Christ our Lord. <em>Amen.</em></br>`,
		},
		Prayer{
			Name: "For the Aged",
			Content: `Look with mercy, O God our Father, on all whose increasing
			years bring them weakness, distress, or isolation. Provide for
			them homes of dignity and peace; give them understanding
			helpers, and the willingness to accept help; and, as their
			strength diminishes, increase their faith and their assurance
			of your love. This we ask in the name of Jesus Christ our
			Lord. <em>Amen.</em>`,
		},
	},
}
