package main

import (
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

func Test__sould(t *testing.T) {
	// {{"Library", 19030}, {"Music", 26533}, {"Music Videos", 33968}, {"TV & Films", 33971}, {"Films", 33977}, {"Home Videos", 33980}, {"TV Programmes", 33983}, {"Podcasts", 33986}, {"iTunes U", 34755}, {"Audiobooks", 34758}, {"Books", 34763}, {"PDFs", 34766}, {"Audiobooks", 34769}, {"Genius", 34804}, {"@music", 34807}, {"_audiobooks", 40808, 34807}, {"Dekkers, Midas", 41504, 40808}, {"De Flamingo en andere beesten", 41508, 41504}, {"Harry Potter", 41512, 40808}, {"2 Harry Potter und die Kammer des Schreckens", 42192, 41512}, {"und die Kammer des Schreckens - CD 01", 42324, 42192}, {"und die Kammer des Schreckens - CD 02", 42341, 42192}, {"und die Kammer des Schreckens - CD 03", 42358, 42192}, {"und die Kammer des Schreckens - CD 04", 42373, 42192}, {"und die Kammer des Schreckens - CD 05", 42389, 42192}, {"und die Kammer des Schreckens - CD 06", 42406, 42192}, {"und die Kammer des Schreckens - CD 07", 42421, 42192}, {"und die Kammer des Schreckens - CD 08", 42437, 42192}, {"und die Kammer des Schreckens - CD 09", 42454, 42192}, {"und die Kammer des Schreckens - CD 10", 42468, 42192}, {"4 Harry Potter und der Feuerkelch", 42483, 41512}, {"01 Das Haus der Riddles", 42754, 42483}, {"02 Die Narbe", 42764, 42483}, {"03 Die Einladung", 42771, 42483}, {"04 Zuruck zum Fuchsbau", 42778, 42483}, {"05 Weasleys Zauberhafte Zauberscherze", 42785, 42483}, {"06 Der Portschlüssel", 42793, 42483}, {"07 Bagman und Crouch", 42799, 42483}, {"08 Die Quidditch-Weltmeisterschaft", 42809, 42483}, {"09 Das dunkle Mal", 42820, 42483}, {"10 Wirbel im Ministerium", 42833, 42483}, {"11 Im Hogwarts Express", 42840, 42483}, {"12 Das Trimagische Turnier", 42847, 42483}, {"13 Mad-Eye Moody", 42859, 42483}, {"14 Die Unverzeihlichen Flüche", 42867, 42483}, {"15 Beauxbatonsund Durmstrang", 42877, 42483}, {"16 Der Feuerkelch", 42887, 42483}, {"17 Die vier Champions", 42899, 42483}, {"18 Die Eichung der Zauberstäbe", 42908, 42483}, {"19 Der Ungarische Hornschwanz", 42920, 42483}, {"20 Die erste Aufgabe", 42932, 42483}, {"21 Die Hauselfen-Befreiungsfront", 42944, 42483}, {"22 Die unerwartete Aufgabe", 42955, 42483}, {"23 Der Weihnachtsball", 42964, 42483}, {"24 Rita Kimmkorns Riesenknüller", 42978, 42483}, {"25 Das Ei und das Auge", 42989, 42483}, {"26 Die zweite Aufgabe", 43000, 42483}, {"27 Tatzes Rückkehr", 43014, 42483}, {"28 Mr Crouchs Wahn", 43027, 42483}, {"29 Der Traum", 43040, 42483}, {"30 Das Denkarium", 43049, 42483}, {"31 Die dritte Aufgabe", 43062, 42483}, {"32 Fleisch, Blut und Knochen", 43076, 42483}, {"33 Die Todesser", 43082, 42483}, {"34 Priori Incantatem", 43092, 42483}, {"35 Veritaserum", 43099, 42483}, {"36 Die Wege trennen sich", 43111, 42483}, {"37 Der Anfang", 43123, 42483}, {"1 Harry Potter und der Stein der Weisen", 43133, 41512}, {"3 Harry Potter und der Gefangene von Askaban", 43244, 41512}, {"6 Harry Potter und der Halbblutprinz", 43395, 41512}, {"Hawking, Stephen", 43422, 40808}, {"A Brief History of Time", 43426, 43422}, {"The Hitchhiker's Guide to the Galaxy", 43430, 40808}, {"New", 43435, 43430}, {"The Hitchhiker's Guide to the Galaxy", 43440, 40808}, {"_comedy", 43455, 34807}, {"Cabaret", 44036, 43455}, {"Harrie Jekkers", 44515, 44036}, {"Het Geheim Van De Lachende Piccolo", 44608, 44515}, {"CD 1", 44635, 44608}, {"CD 2", 44649, 44608}, {"Met Een Goudvis Naar Zee", 44665, 44515}, {"CD 1", 44701, 44665}, {"CD 2", 44721, 44665}, {"De Man In De Wolken", 44740, 44515}, {"Het Gelijk Van De Koffietent", 44764, 44515}, {"Herman Finkers", 44779, 44036}, {"1995 Geen spatader veranderd", 45036, 44779}, {"CD1", 45074, 45036}, {"CD2", 45096, 45036}, {"1999 Zijn minst beroerde liedjes", 45115, 44779}, {"Cd1 Tweants", 45177, 45115}, {"CD2 Zo nederlands mogelijk", 45210, 45115}, {"2002 Kalm Aan En Rap Een Beetje", 45242, 44779}, {"CD1", 45283, 45242}, {"CD2", 45304, 45242}, {"1983 Als gezonde jongen zijnde", 45327, 44779}, {"1985 EHBO is mijn lust en mijn leven", 45345, 44779}, {"1987 Het meisje van de slijterij", 45360, 44779}, {"1987 Van zijn elpee", 45380, 44779}, {"1990 De zon gaat zinloos onder - Morgen moet ze toch weer op", 45396, 44779}, {"1992 Dat heeft zo'n jongen toch niet nodig (het meisje met de eierstokjes)", 45412, 44779}, {"1993 Het Carnaval der Dieren", 45434, 44779}, {"mannen", 45470, 44036}, {"mannen1", 45499, 45470}, {"mannen2", 45515, 45470}, {"Grolsch stand up & comedy live", 45531, 44036}, {"Hans Teeuwen Industry Of Love", 45543, 44036}, {"Jerry Seinfeld", 45571, 44036}, {"Lebbis & Jansen - Jakkeren door 1998", 45595, 44036}, {"mannen", 45610, 44036}, {"Helge Schneider - Aprikose, Banane, Erdbeer", 45652, 43455}, {"CD 1", 45710, 45652}, {"CD 2", 45732, 45652}, {"CD 3", 45750, 45652}, {"XFM - Ricky Gervais Show", 45774, 43455}, {"series2", 45796, 45774}, {"Cabaret", 45818, 43455}, {"Comedy Club Shootout Vol. 1 -Classic Stand-Up with Jerry Seinfeld Bill Maher David Spade and many more", 45825, 43455}, {"XFM - Ricky Gervais Show", 45829, 43455}, {"_style", 45855, 34807}, {"AcidFlash", 45863, 45855}, {"Acid Flash vol. 1 [Disk 1]", 45871, 45863}, {"Adele", 45879,34807}, {"Adele - 19", 45916, 45879}, {"Adele - 21", 45936, 45879}, {"Air", 45956, 34807}, {"Air - 10000 Hz Legend", 45990, 45956}, {"Air - Moon Safari (1998)", 46004, 45956}, {"Air - Talkie Walkie", 46017, 45956}, {"Alicia Keys", 46030, 34807}, {"Alicia Keys - As I Am", 46047, 46030}, {"Amon Tobin", 46064, 34807}, {"amon tobin - permutation", 46091, 46064}, {"amon tobin - supermodified", 46106, 46064}, {"Amsterdam Klezmer Band", 46121, 34807}, {"amsterdam klezmer band - katakofti", 46168, 46121}, {"amsterdam klezmer band - limonchiki", 46185, 46121}, {"Amsterdam Klezmer Band - Man Bites Dog Eats", 46198, 46121}, {"Amsterdam Klezmer Band - Remixed", 46208, 46121}, {"Amy Winehouse", 46224, 34807}, {"Amy Winehouse - Back To Black", 46258, 46224}, {"Amy Winehouse - Back To Black (Live at Paradiso)", 46271, 46224}, {"Amy Winehouse - Frank", 46280, 46224}, {"Ane Brun", 46298, 34807}, {"Ane Brun - Spending Time With Morgan", 46313, 46298}, {"Aphrodite", 46328, 34807}, {"Aphrodite - Urban Jungle", 46347, 46328}, {"Arctic Monkeys", 46366, 34807}, {"Arctic Monkeys - Artic Monkeys (Album)", 46397, 46366}, {"Arctic Monkeys - Favourite Worst Nightmare", 46416, 46366}, {"Arts The Beatdoctor", 46431, 34807}, {"Arts The Beatdoctor - Transitions (2007)", 46449, 46431}, {"Ascii Disco", 46467, 34807}, {"ascii disko - ascii disko", 46482, 46467}, {"Audio Bullys", 46497, 34807}, {"Audio Bullys - Ego War", 46514, 46497}, {"Basement Jaxx", 46531, 34807}, {"Basement Jaxx - Crazy Itch Radio", 46542, 46531}, {"Black Eyed Peas", 46553, 34807}, {"Black Eyed Peas - E.N.D", 46596, 46553}, {"CD 1", 46624, 46596}, {"CD 2", 46642, 46596}, {"Black Eyed Peas - Monkey Business (2005)", 46655, 46553}, {"Bloc Party", 46673, 34807}, {"Bloc Party - A Weekend In The City", 46700,46673}, {"Bloc Party - Silent Alarm", 46714, 46673}, {"Bloodhound Gang", 46730, 34807}, {"Bloodhound Gang - Hefty Fine (2005)", 46745, 46730}, {"Blueprint", 46760, 34807}, {"blueprint (disk 1) - various", 46798, 46760}, {"blueprint (disk 2) - various",46812, 46760}, {"blueprint (disk 3) - various", 46825, 46760}, {"Bob Marley", 46842, 34807}, {"Bob Marley - Various albums", 46858, 46842}, {"Bonobo", 46874, 34807}, {"Bonobo - Animal Magic", 46887, 46874}, {"C-mon & Kypski", 46900, 34807}, {"C-mon & Kypski - Live at the Dom", 46943, 46900}, {"C-Mon & Kypski - Static Traveller", 46952, 46900}, {"C-Mon & Kypski - Vinyl Voodoo", 46966, 46900}, {"C-mon & Kypski - Where the wild things are", 46980, 46900}, {"Camille", 46995, 34807}, {"Camille - le fil", 47013, 46995}, {"Cassius", 47031, 34807}, {"Cassius - Au Reve", 47048, 47031}, {"Chemical Brothers by ViSenT", 47065, 34807}, {"4 Singles", 47148, 47065}, {"The Chemical Brothers - Come With Us_The test", 47163, 47148}, {"The Chemical Brothers - Hey BoyHey Girl", 47169, 47148}, {"The Chemical Brothers - It Began In Africa", 47175, 47148}, {"The Chemical Brothers - Out of Control", 47181, 47148}, {"Chemical Brothers - Brothers gonna work it out", 47187, 47065}, {"Chemical Brothers - Dig Your Own Hole", 47213, 47065}, {"Chemical Brothers - Surrender", 47227, 47065}, {"Chemical Brothers- Come whit us", 47241, 47065}, {"Chemical Brothers- exit planet dust", 47254, 47065}, {"Extra", 47268, 47065}, {"Coldplay", 47273, 34807}, {"Coldplay - A rush of blood", 47319, 47273}, {"Coldplay - EP's ed", 47333, 47273}, {"Coldplay - Parachutes (2000)", 47345, 47273}, {"Coldplay - X&Y (2000)", 47358, 47273}, {"Coparck", 47374, 34807}, {"Birds, Happiness & Still not Worried", 47417, 47374}, {"Few Chances Come Once In A Lifetime", 47435, 47374}, {"The 3rd Album", 47451, 47374}, {"Cowboy Bebop", 47466, 34807}, {"Yoko Kanno - Cowboy Bebop All Albums Ost", 47558, 47466}, {"Cowboy Bebop - Ask DNA", 47650, 47558}, {"Cowboy Bebop - Cowgirl Ed OST", 47658, 47558}, {"Cowboy Bebop - Future Blues OST", 47667, 47558}, {"Cowboy Bebop - OST1", 47688, 47558}, {"Cowboy Bebop - OST2 - No Disc", 47708, 47558}, {"Cowboy Bebop - OST3 - Blue", 47729, 47558}, {"Cowboy Bebop - Vitamineless", 47749, 47558}, {"Crystal Fighters", 47760, 34807}, {"Crystal Fighters - Star Of Love", 47774, 47760}, {"Daft Punk", 47788, 34807}, {"Daft Punk - Human After All (2005)", 47801, 47788}, {"Die Ärzte", 47814, 34807}, {"2007 - Jazz ist anders", 47855, 47814}, {"Bonus EP", 47861, 47855}, {"2007 - Jazz ist anders", 47867, 47814}, {"Rock n Roll Realschule", 47886, 47814}, {"Die fantastischen Vier", 47908, 34807}, {"Die fantastischen Vier - MTV unplugged", 47927, 47908}, {"Digitalism", 47946, 34807}, {"Digitalism - Digitalism", 47964, 47946}, {"Dire Straits", 47982, 34807}, {"Dire Straits - Sultans Of Swing 'The Very Best Of'", 48017, 47982}, {"Cd1", 48043, 48017}, {"Cd2", 48062, 48017}, {"Dire Straits - Brothers In Arms", 48072, 47982}, {"Doe Maar", 48084, 34807}, {"Doe Maar - Echt Alles (3 CD Verzamelbox)", 48167, 48084}, {"CD 1", 48225, 48167}, {"CD 2", 48247, 48167}, {"CD 3", 48268, 48167}, {"Doe Maar - Doe de dub", 48289, 48084}, {"Doe Maar - KLAAR", 48299, 48084}, {"Editors", 48320, 34807}, {"Editors - The Back Room (2005) + 2 cd special edition 2007", 48340, 48320}, {"CUTTINGS disc2", 48349, 48340}, {"Editors - The Back Room (2005) + 2 cd special edition 2007", 48358, 48320}, {"Eric Clapton", 48372, 34807}, {"Eric Clapton - Unplugged", 48389, 48372}, {"Faithless", 48406, 34807}, {"Faithless - Forever Faithless greatest hits (2005)", 48440, 48406}, {"Faithless - No Roots (2005)", 48459, 48406}, {"Fatboy Slim", 48477, 34807}, {"Fatboy Slim - Halfway between the gutter and the stars", 48514, 48477}, {"Fatboy Slim - Palookaville", 48528, 48477}, {"Fatboy Slim - You've come a long way", 48543, 48477}, {"Fink", 48557, 34807}, {"Fink - Sort of Revolution (2009) [VIPER666]", 48570, 48557}, {"Flight of the Conchords", 48583, 34807}, {"Flight Of The Conchords - All Songs Season 1-2", 48699, 48583}, {"Flight of the Conchords - Folk The World Tour", 48736, 48583}, {"Flight of the Conchords - HBO One Night Stand", 48753, 48583}, {"Flight of the Conchords - Other Flight of the Conchords", 48768, 48583}, {"Flight of the Conchords - Season 1 Music", 48773, 48583}, {"Flight of the Conchords - The Distant Future", 48802, 48583}, {"Flight of the Conchords - World Comedy 2004", 48811, 48583}, {"Flight of the Conchords (April, 2008)", 48818, 48583}, {"Forever Classics", 48836, 34807}, {"101 Classical Greats Cd1", 49102, 48836}, {"101 Classical Greats Cd2", 49126, 48836}, {"101 Classical Greats Cd3", 49149, 48836}, {"101 Classical Greats Cd4", 49172, 48836}, {"Bach", 49195, 48836}, {"Bahms", 49209, 48836}, {"Beethoven", 49220, 48836}, {"Chopin", 49231, 48836}, {"Dvorak", 49242, 48836}, {"Grieg", 49254, 48836}, {"Händel", 49268, 48836}, {"Haydn", 49299, 48836}, {"Mozart", 49314, 48836}, {"Mussorgsky", 49327, 48836}, {"Ravel", 49348, 48836}, {"Schumann", 49365, 48836}, {"Strauss", 49378, 48836}, {"Tchaikovsky", 49389, 48836}, {"Vivaldi", 49398, 48836}, {"Franz Ferdinand", 49422, 34807}, {"Franz Ferdinand - Franz Ferdinand", 49449, 49422}, {"Franz Ferdinand - You Could Have It So Much Better With Franz Ferdinand (2005)", 49463, 49422}, {"Freestylers", 49479, 34807}, {"Freestylers - Raw As Fuck", 49495, 49479}, {"Goose", 49511, 34807}, {"Goose - Bring It On", 49526, 49511}, {"Gorillaz", 49541, 34807}, {"Gorillaz - Demon Days (2005)", 49574, 49541}, {"Gorillaz - Gorillaz (2001)", 49592, 49541}, {"Greenday", 49610, 34807}, {"Greenday - American Idiot", 49655, 49610}, {"Greenday - Dookie", 49671, 49610}, {"Greenday - Insomniac", 49689, 49610}, {"Helge Schneider", 49706, 34807}, {"Helge Schneider - I Brake Together (2007) - Pop [www.torrentazos.com]", 49763, 49706}, {"Helge Schneider- The Best Of", 49781, 49706}, {"Hollandse Hits", 49823, 34807}, {"Guus Meeuwis - [2005] 10 Jaar - Levensecht", 49841, 49823}, {"Holy Fuck", 49859, 34807}, {"Holy Fuck-LP", 49871, 49859}, {"Jack Johnson", 49883, 34807},{"Jack Johnson - Brushfire Fairytales", 49967, 49883}, {"Jack Johnson - In Between Dreams", 49983, 49883}, {"Jack Johnson - Live And Acoustic At Kfog", 50000, 49883}, {"Jack Johnson - On And On", 50014, 49883}, {"Jack Johnson - Thicker Than Water Soundtrack", 50033, 49883}, {"Sing-A-Longs And Lullabies For The Film Curious George", 50050, 49883}, {"Jamie Cullum", 50066, 34807}, {"Jamie Cullum - Catching Tales (2005)", 50101, 50066}, {"Jamie Cullum - Twentysomething Special Edition (2004)", 50118, 50066}, {"Jamiroquai", 50139, 34807}, {"00-Jamiroquai - Emergency on Planet Earth (1993) 192kbps", 50252, 50139}, {"01-Jamiroquai - The Return of the Space Cowboy (1994) 192kbps", 50265, 50139}, {"02-Jamiroquai - Travelling Without Moving (1996) 192kbps", 50279, 50139}, {"03-Jamiroquai - Synkronized (1999) 192kbps", 50296, 50139}, {"04-Jamiroquai - A Funk Odyssey (2001) 192kbps", 50310, 50139}, {"05-Jamiroquai - Late Night Tales (2003) 192kbps", 50323, 50139}, {"Jamiroquai - Dynamite", 50344, 50139}, {"Jamiroquai - Greatest hits", 50359, 50139}, {"Jamiroquai-Rock_Dust_Light_Star-2010-DOH", 50374, 50139}, {"Jan Delay", 50389, 34807}, {"Jan Delay - Searching for the Jan Soul Rebels (2001) [MP3.320]", 50404, 50389}, {"Jem", 50419, 34807}, {"Jem - Finally Woken", 50433, 50419}, {"Jesus Christ Superstar", 50447, 34807}, {"CD 1", 50476, 50447}, {"CD 2", 50493, 50447}, {"Jimi Hendrix", 50508, 34807}, {"Jimi Hendrix Experience - Are You Experienced", 50549, 50508}, {"Jimi Hendrix Experience - Axis Bold as love", 50563, 50508}, {"The Best of", 50579, 50508}, {"JJ Cale", 50596, 34807}, {"JJ Cale - Okie", 50631, 50596}, {"JJ Cale - The very best of", 50646, 50596}, {"Johnnhy Cash", 50669, 34807}, {"2003 - Unearthed", 50807, 50669}, {"CD1", 50889, 50807}, {"CD2", 50910, 50807}, {"CD3", 50929, 50807}, {"CD4", 50947, 50807}, {"CD5", 50965, 50807}, {"1994 - American Recordings", 50983, 50669}, {"1996 - American II - Unchained", 50999, 50669}, {"2000 - American III - Solitary Man", 51016, 50669}, {"2002 - American IV - The Man Comes Around", 51033, 50669}, {"Jojo Mayer´s nerve", 51051, 34807}, {"Jojo Mayer & Nerve - Live @ Modern Drummer Festival 2005 (320 Kbps)", 51067, 51051}, {"Jojo Mayer & Nerve - Prohibited Beats (2007)", 51074, 51051}, {"Justice", 51086,34807}, {"Justice - Cross (2008)", 51102, 51086}, {"Kashmir", 51118, 34807}, {"Cruzential", 51188, 51118}, {"Kashmir, live fra taget af Radiohuset", 51202, 51118}, {"No Balance Palace", 51216, 51118}, {"The Good Life", 51230, 51118}, {"Travelogue", 51243, 51118}, {"Zitilites", 51256, 51118}, {"Keane", 51273, 34807}, {"Keane - Hopes And Fears (2004)", 51287, 51273}, {"Keith Jarrett", 51301, 34807}, {"Keith Jarrett - The Koln Concert", 51308, 51301}, {"Kraak & Smaak", 51315, 34807}, {"Kraak And Smaak -Plastic People", 51331, 51315}, {"Kraftwerk", 51347, 34807}, {"Kraftwerk - Computer World", 51398, 51347}, {"Kraftwerk - Downloads", 51408, 51347}, {"Kraftwerk - Man Machine", 51416, 51347}, {"Kraftwerk - Radioactivity", 51425, 51347}, {"Kraftwerk - The Mix", 51440, 51347}, {"Kraftwerk - Trans Europe Express", 51454, 51347}, {"Kyteman", 51464, 34807}, {"Kyteman - The Hermit Sessions", 51481, 51464}, {"Lenny Kravitz", 51498, 34807}, {"Lenny Kravitz - Greatest Hits [2000]", 51530, 51498}, {"Lenny Kravitz - Mama Said [1991]", 51548, 51498}, {"Lucky Fonz III", 51565, 34807}, {"Lucky Fonz III - Life is short", 51594, 51565}, {"Lucky Fonz III - Lucky Fonz III", 51611, 51565}, {"Madonna", 51626, 34807}, {"Madonna - Confessions On A Dance Floor (2005)", 51641, 51626}, {"Mark Ronson", 51656, 34807}, {"Mark Ronson - 2007 Version", 51673, 51656}, {"Massive Attack", 51690, 34807}, {"Massive Attack - 100th Window", 51755, 51690}, {"Massive attack - Blue Lines", 51768, 51690}, {"Massive Attack - Mezzanine", 51789, 51690}, {"Massive Attack - Nexus - Bootleg", 51803, 51690}, {"Massive Attack - Protection", 51819, 51690}, {"Melody Gardot", 51832, 34807}, {"Melody Gardot - My One And Only Thrill (2009)", 51857, 51832}, {"Melody Gardot - Worrisome Heart (2008)", 51872, 51832}, {"Michael Jackson", 51885, 34807}, {"Michael Jackson - Number Ones", 51906, 51885}, {"Mika", 51927, 34807}, {"Mika - Life In Cartoon Motion (2007)", 51942, 51927}, {"Moby", 51957, 34807}, {"Moby - Hotel (2005)", 51974, 51957}, {"Mr. Oizo", 51991, 34807}, {"Mr. Oizo - Flat beat", 52005, 51991}, {"Mumford and Sons", 52019, 34807}, {"Mumford and Sons - Sigh No More (2009)", 52034, 52019}, {"Muse", 52049, 34807}, {"Muse - 2002 - Microcuts On Stage", 52112, 52049}, {"Muse - 2003 - Absolution", 52132, 52049}, {"Muse - 2006 - Black Holes and Revelations", 52149, 52049}, {"Muse - Misc", 52163, 52049}, {"N.E.R.D", 52184, 34807}, {"N.E.R.D - Fly Or Die", 52226, 52184}, {"N.E.R.D - In Search Of", 52241, 52184}, {"N.E.R.D - Seeing Sounds", 52256,52184}, {"Nizlopi", 52274, 34807}, {"Nizlopi - Half These Songs Are About You [2004]", 52288, 52274}, {"No Doubt", 52302, 34807}, {"No Doubt - No Doubt", 52332, 52302}, {"No Doubt - Rock Steady", 52349, 52302}, {"Ohrbooten", 52365, 34807}, {"Ohrbooten - Babylon bei Boot", 52395, 52365}, {"Ohrbooten - Spieltrieb", 52412, 52365}, {"Paolo Conte", 52428, 34807}, {"Paolo Conte - The Best Of (1998)", 52451, 52428}, {"Pendulum", 52474, 34807}, {"Pendulum - 2008 - In Silico", 52501, 52474}, {"Pendulum - HoldYour Colour", 52514, 52474}, {"Pete Philly & Perquisite", 52531, 34807}, {"Pete Philly & Perquisite - Mindstate", 52567, 52531}, {"Pete Philly & Perquisite - Mystery Repeats", 52587, 52531}, {"Peter Fox", 52606, 34807}, {"Peter Fox - Stadtaffe", 52621,52606}, {"Pink", 52636, 34807}, {"Pink - Greatest Hits So Far (2010)", 52660, 52636}, {"Pop Classics", 52684, 34807}, {"Pop Classics - The long versions (cd1)", 52707, 52684}, {"Pop Classics - The Long versions (cd2)", 52720, 52684}, {"Queens of the Stoneage", 52733, 34807}, {"Queens of the Stoneage - Songs for the deaf", 52751, 52733}, {"Radiohead", 52769, 34807}, {"Radiohead - In Rainbows (2007)", 52871, 52769}, {"CD1", 52892, 52871}, {"CD2", 52905, 52871}, {"Radiohead - Amnesiac (2001)", 52916, 52769}, {"Radiohead - Hail to the Thief (2003)", 52930, 52769}, {"Radiohead - In Rainbows From the Basement", 52947, 52769}, {"Radiohead - Kid A (2000)", 52959, 52769}, {"Radiohead - OK Computer (1997)", 52972, 52769}, {"Radiohead - Pablo Honey (1993)", 52987, 52769}, {"Radiohead - The Bends (1995)", 53003, 52769}, {"Ramblers", 53018, 34807}, {"Ramblers Collectie", 53041, 53018}, {"Ray Charles", 53064, 34807}, {"His Greatest Hits CD1", 53139, 53064}, {"His Greatest Hits CD2", 53165, 53064}, {"Ray Charles -  Genius & Friends - 2005", 53188, 53064}, {"Ray Charles - Blues And Jazz", 53205, 53064}, {"Ray Charles - Genius Loves Company (2004)", 53210, 53064}, {"Red Hot Chili Peppers", 53226, 34807}, {"Red Hot Chili Peppers - Blood Sugar Sex Magik", 53314, 53226}, {"Red Hot Chili Peppers - Californication", 53334, 53226}, {"Red Hot Chili Peppers - Mothers Milk", 53352, 53226}, {"Red Hot Chili Peppers - One Hot Minute", 53368, 53226}, {"Red Hot Chili Peppers - Stadium Arcadium", 53384, 53226}, {"Regina Spektor", 53414, 34807}, {"Regina Spektor - Begin To Hope", 53477, 53414}, {"Regina Spektor - Eleven Eleven", 53492, 53414}, {"Regina Spektor - Far", 53507, 53414}, {"Regina Spektor - Songs", 53523, 53414}, {"Regina Spektor - Soviet Kitsch", 53538, 53414},{"Richard Cheese", 53552, 34807}, {"Richard Cheese - Aperitif for Destruction (2005)", 53624, 53552}, {"Richard Cheese - I'd like a virgin (2006)", 53643, 53552}, {"Richard Cheese - Lounge against the Machine (2000)", 53665, 53552}, {"Richard Cheese - Tuxicity (2003)", 53684, 53552}, {"RJD2", 53705, 34807}, {"2002 - Deadringer", 53739, 53705}, {"2002 - DeadSampler", 53758, 53705}, {"Rob Costlow", 53776, 34807}, {"Rob Costlow - Contemporary Piano - Sophomore Jinx", 53798, 53776}, {"Rob Costlow - Contemporary Piano - Woods of Chaos", 53810, 53776}, {"Robbie Williams", 53823, 34807}, {"Robbie Williams - Intensive Care", 53838, 53823}, {"Röyksopp", 53853, 34807}, {"Röyksopp Discography", 54001, 53853}, {"49 Percent Remixes", 54124, 54001}, {"Beautiful Day Without You Remixes", 54133, 54001}, {"Eple Remixes", 54139, 54001}, {"Junior", 54147, 54001}, {"Melody AM", 54161, 54001}, {"Miscellaneous and Rarities", 54174, 54001}, {"Only This Moment Remixes", 54186, 54001}, {"Poor Leno Remixes", 54194, 54001}, {"Remind Me Remixes", 54204, 54001}, {"Remixes by Royksopp", 54212, 54001}, {"Röyksopp's Night Out", 54234, 54001}, {"So Easy Remixes", 54246, 54001}, {"Sparks Remixes", 54250, 54001}, {"The Understanding - Disc 1", 54259, 54001}, {"The Understanding - Disc 2 (Deluxe Edition)", 54274, 54001}, {"What Else is There Remixes", 54282, 54001}, {"Röyksopp - Junior (2009)", 54292, 53853}, {"Royksopp - Melody AM (2002)", 54307, 53853}, {"Saint Germain", 54323, 34807}, {"Saint Germain - Boulevard", 54343, 54323}, {"Saint Germain - Tourist", 54354, 54323}, {"Samy Deluxe", 54366, 34807}, {"Samy Deluxe - Schwarzweiss (2011) - Hip-Hop [www.torrentazos.com]", 54389, 54366}, {"Bonus CD", 54396, 54389}, {"Samy Deluxe - Schwarzweiss (2011) - Hip-Hop [www.torrentazos.com]", 54403, 54366}, {"Seeed", 54422, 34807}, {"2012 Seeed", 54475, 54422}, {"Seeed - Music Monks", 54489, 54422}, {"Seeed - New Dubby Conquerors", 54505, 54422}, {"Seeed - Next!", 54520, 54422}, {"Snow Patrol", 54537, 34807}, {"Snow Patrol - Final Straw", 54554, 54537}, {"Soulwax", 54571, 34807}, {"Soulwax - Any Minute Now", 54587, 54571}, {"Spinvis", 54603, 34807}, {"Spinvis - Dagen van Gras, Dagen van Stro", 54640, 54603}, {"Spinvis - Herfst en Nieuwegein", 54654, 54603}, {"Spinvis - Spinvis", 54667, 54603}, {"Stereo Total", 54683, 34807}, {"02 - Monokin", 54784, 54683}, {"03 - Juke Box Alarm", 54802, 54683}, {"03 - My Melody", 54821, 54683}, {"04 - Musique Automatique", 54840, 54683}, {"05 - Do The Bambi", 54859, 54683}, {"06 - Discotheque", 54881, 54683}, {"Stromae", 54900, 34807}, {"Stromae - Cheese", 54914, 54900}, {"Tenacious D", 54928, 34807}, {"Overig", 55008, 54928}, {"Tenacious D - TenaciousD", 55036, 54928}, {"Tenacious D - The Complete Masterworks", 55060, 54928}, {"Tenacious D - The Pick Of Destiny", 55079, 54928}, {"The Doors", 55097, 34807}, {"Friday April 10 At Boston Arena", 55611, 55097}, {"The Doors - Friday April 10 At Boston Arena (1st Show)", 55660, 55611}, {"The Doors - Friday April 10 At Boston Arena (2st Show)", 55679, 55611}, {"Set The Night On Fire- The Doors Bright Midnight Archives Concerts (Live)", 55712, 55097}, {"CD1", 55796, 55712}, {"CD2", 55826, 55712}, {"CD3", 55850, 55712}, {"The Doors - STUDIO DISCOGRAPHY", 55886, 55097}, {"The Doors - 1978 - An American Prayer", 55912, 55886}, {"Absolutely Live", 55938, 55097}, {"Backstage & Dangerous- The Private Rehearsal", 55962, 55097}, {"Box Set", 55989, 55097}, {"Essential Rarities", 56045, 55097}, {"L.A. Woman (40th Anniversary Mixes)", 56063, 55097}, {"LiveAt The Aquarius Theater- The First Performance", 56078, 55097}, {"Live At The Aquarius Theater- The Second Performance", 56102, 55097}, {"Live In Detroit", 56138, 55097}, {"Live In Hollywood- Highlights From The Aquarius Theatre Performances", 56168, 55097}, {"Morrison Hotel (40th Anniversary Mixes)", 56181, 55097}, {"Strange Days (40th Anniversary Mixes)", 56205, 55097}, {"The Bright Midnight Sampler (Live)", 56220, 55097}, {"The Doors - An American Prayer", 56236, 55097}, {"The Doors - The Best Of The Doors SPECIAL EDITION (2002)", 56240, 55097}, {"The Doors - The Complete Matrix Club Tapes Box Set (cd1)", 56260, 55097}, {"The Doors - The Complete Matrix Club Tapes Box Set (cd2)", 56275, 55097}, {"The Doors - The Complete Matrix Club Tapes Box Set (cd3)", 56286, 55097}, {"The Doors - The Complete Matrix Club Tapes Box Set (cd4)", 56297, 55097}, {"The Doors (40th Anniversary Mixes)", 56308, 55097}, {"The Soft Parade (40th Anniversary Mixes)", 56325, 55097}, {"Waiting For The Sun (40th Anniversary Mixes)", 56343, 55097}, {"The Infadels", 56362, 34807}, {"The Infadels - We are not the infadels", 56376, 56362}, {"The Presidents of the USA", 56390, 34807}, {"Freaked out and small", 56448, 56390}, {"Presidents of the USA - These Are The Good Times People", 56463, 56390}, {"The Presidents of the USA  Pt. 1", 56480, 56390}, {"The Presidents of the USA Pt.1", 56484, 56390}, {"The Presidents of the USA Pt.2", 56500, 56390}, {"The Prodigy", 56518, 34807}, {"The Prodigy - 1997 - The Fat Of The Land", 56554, 56518}, {"The Prodigy - 2004 - Always Outnumbered Never Outgunned", 56567, 56518}, {"The Prodigy - Invaders must die", 56582, 56518}, {"The Specials", 56596, 34807}, {"The Specials", 56614, 56596}, {"Tryo", 56632, 34807}, {"Tryo - Grain de sable",56649, 56632}, {"Unkle", 56666, 34807}, {"Unkle - Psyence Fiction", 56696, 56666}, {"UNKLE - War Stories [Mp3-128-2007]", 56712, 56666}, {"Urban Dance Squad", 56729, 34807}, {"Urban Dance Squad - Mental Floss For The Globe", 56745, 56729}, {"Xploding Plastix", 56761, 34807}, {"Xploding Plastix - The Donca Matic Singalongs", 56789, 56761}, {"Xploding Plastix - Xploding Plastix", 56803, 56761}, {"Yann Tiersen", 56820, 34807}, {"Amelie", 56866, 56820}, {"Good bye lenin", 56889, 56820}, {"_audiobooks", 56915, 34807}, {"_classic", 56920, 34807}, {"_losse nummers", 56925, 34807}, {"_mixes", 56929, 34807}, {"~Audioboeken", 56945, 34807}, {"Amsterdam Klezmer Band", 56948, 34807}, {"Black Eyed Peas", 56952, 34807}, {"Command & Conquer - Soundtrack Collection", 56956, 34807}, {"Jan Delay", 57028, 34807}, {"Neil Diamond - Home before Dark", 57032, 34807}, {"Samy Deluxe", 57047, 34807}, {"Sound of Music", 57051, 34807}, {"The Chap", 57070, 34807}, {"The Fifth Element (Eric Serra, 1997)", 57079, 34807}, {"Total Annihilation - Game Soundtrack (1997)", 57108, 34807}, {"Traditional Japanese Music - Kodo", 57127, 34807}, {"Voice Memos", 57180}}

	assert.Equal(t, "\"@music\", 34807", parse(`{"@music", 34807}`))
}

func Test__should_lex_root_item(t *testing.T) {
	_, items := lex("name", "{\"bla\", 19}")

	assert.Equal(t, item{plainText, ""}, <-items)
	assert.Equal(t, item{opener, "{"}, <-items)
	assert.Equal(t, item{openQuote, "\""}, <-items)
	assert.Equal(t, item{itemName, "bla"}, <-items)
	assert.Equal(t, item{closeQuote, "\""}, <-items)
	assert.Equal(t, item{separator, ", "}, <-items)
	assert.Equal(t, item{itemId, "19"}, <-items)
	assert.Equal(t, item{closer, "}"}, <-items)
}

func Test__should_lex_complete_item(t *testing.T) {
	_, items := lex("name", "{\"bla\", 19, 20}")

	assert.Equal(t, item{plainText, ""}, <-items)
	assert.Equal(t, item{opener, "{"}, <-items)
	assert.Equal(t, item{openQuote, "\""}, <-items)
	assert.Equal(t, item{itemName, "bla"}, <-items)
	assert.Equal(t, item{closeQuote, "\""}, <-items)
	assert.Equal(t, item{separator, ", "}, <-items)
	assert.Equal(t, item{itemId, "19"}, <-items)
	assert.Equal(t, item{separator, ", "}, <-items)
	assert.Equal(t, item{itemId, "20"}, <-items)
	assert.Equal(t, item{closer, "}"}, <-items)
}

func Test__should_lex_multiple_items(t *testing.T) {
	_, items := lex("name", "{{\"bla\"}, {\"B\", 19}}")

	assert.Equal(t, item{plainText, ""}, <-items)
	assert.Equal(t, item{opener, "{"}, <-items)
	assert.Equal(t, item{opener, "{"}, <-items)
	assert.Equal(t, item{openQuote, "\""}, <-items)
	assert.Equal(t, item{itemName, "bla"}, <-items)
	assert.Equal(t, item{closeQuote, "\""}, <-items)
	assert.Equal(t, item{closer, "}"}, <-items)
	assert.Equal(t, item{separator, ", "}, <-items)
	assert.Equal(t, item{opener, "{"}, <-items)
	assert.Equal(t, item{openQuote, "\""}, <-items)
	assert.Equal(t, item{itemName, "B"}, <-items)
	assert.Equal(t, item{closeQuote, "\""}, <-items)
	assert.Equal(t, item{separator, ", "}, <-items)
	assert.Equal(t, item{itemId, "19"}, <-items)
	assert.Equal(t, item{closer, "}"}, <-items)
	assert.Equal(t, item{closer, "}"}, <-items)
}

type stateFn func(*lexer) stateFn

type item struct {
	typ itemType //The type of this item.
	val string   // The value of this item.
}

const (
	itemError  itemType = iota // error occured. Value is error text
	plainText                  //
	opener                     // {
	openQuote                  // "
	closeQuote                 // "
	closer                     // }
	itemName                   //name of the playlist
	itemId                     //name of the playlist
	separator                  // ',' with or without ' '
)

type lexer struct {
	name  string    // used only for Error reports.
	input string    // the strig being scanned
	start int       // start position of this item
	pos   int       // current position in the input
	width int       // width of last rune read
	items chan item // channel of scanned items
}

func lex(name, input string) (*lexer, chan item) {
	l := &lexer{
		name:  name,
		input: input,
		items: make(chan item),
	}
	go l.run() // Concurrently run state machine
	return l, l.items
}

func (l *lexer) run() {
	for state := lexText; state != nil; {
		state = state(l)
	}
	close(l.items)
}

func peekNext(l *lexer) string {
	return string(l.input[l.pos])
}

// emit passes an item back to the client.
func (l *lexer) emit(t itemType) {
	l.items <- item{t, l.input[l.start:l.pos]}
	l.start = l.pos
}

func lexText(l *lexer) stateFn {
	for {
		if peekNext(l) == "{" {
			l.emit(plainText)
			return lexOpener
		}
		l.pos++
	}
	return nil
}

func lexOpener(l *lexer) stateFn {
	l.pos++
	l.emit(opener)
	return lexInsideObject
}

func lexCloser(l *lexer) stateFn {
	l.pos++
	l.emit(closer)
	return lexInsideObject
}

func lexInsideObject(l *lexer) stateFn {
	if l.pos < len(l.input) {
		switch peekNext(l) {
		case "{":
			return lexOpener
		case "}":
			return lexCloser
		case "\"":
			return lexOpenQuote
		case ",":
			return lexSeparator
		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
			return lexId
		}
	}
	return nil
}

func lexOpenQuote(l *lexer) stateFn {
	l.pos++
	l.emit(openQuote)
	return lexInsideString
}

func lexCloseQuote(l *lexer) stateFn {
	l.pos++
	l.emit(closeQuote)
	return lexInsideObject
}

func lexInsideString(l *lexer) stateFn {
	for {
		switch peekNext(l) {
		case "\"":
			l.emit(itemName)
			return lexCloseQuote
		}
		l.pos++
	}
	return nil
}

func lexSeparator(l *lexer) stateFn {
	for {
		if peekNext(l) != " " && peekNext(l) != "," {
			// l.pos++ //we want to include the current rune
			l.emit(separator)
			return lexInsideObject
		}
		l.pos++
	}
	return nil
}

func lexId(l *lexer) stateFn {
	for {
		if !unicode.IsDigit([]rune(peekNext(l))[0]) {
			l.emit(itemId)
			return lexInsideObject
		}
		l.pos++
	}
}

var iStart int
var iCur int

func parse(s string) string {
	for i, r := range s {
		if r == '{' {
			iStart = i + 1
			iCur = i
		} else {
			if r != '}' {
				iCur++
			}
		}
	}
	return s[iStart : iCur+1]
}
