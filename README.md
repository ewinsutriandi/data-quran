# Al-Quran Data Repository

This repository contains all data that is related with Al-Quran, from the text, translation, explanation
(tafseer), metadata and font. It's collected from several sources and stored as JSON (except for font,
obviously) to make it easy to use in various programming language.

## Metadata

Metadata is taken from [Tanzil][tanzil-meta]. There are several metadatas available:

- Surah, the chapter of the Quran. There are 114 surahs in the Quran, each divided into ayahs (verses).
- Juz, the division of Quran into 30 parts.
- Hizb, the division of Quran where one juz has two hizb. The hizb itself sometimes is separated into four 
  quarters, so in Quran there are 240 hizb quarters.
- Ruku basically is a paragraph of the Quran. Used to denote a group of thematically related verses in the
  Quran, so that the reciters could identify when to make ruku in Salah without breaking an ongoing topic
  in the Quranic text.
- Page is division of Quran page based on Medina Mushaf.

## Text

There are several type of text that available in this database:

- `simple`, Quran text in modern Arabic writing style (Imla'ei script) which is commonly used nowadays.
- `simple-plain`, like `simple` but without special demonstration of Ikhfas and Idghams.
- `simple-min`, like `simple` but with a minimal number of diacritics and symbols. Suitable for
   embedding in other texts.
- `simple-clean`, like `simple` but without any diacritics or symbols. Suitable for searching.
- `uthmani`, Quran text in old-fashion script which used in Medina Mushaf.
- `uthmani-min`, like `uthmani` but with a minimal number of diacritics and symbols.
- `indonesia`, Quran text that used as a standard for writing Quran in Indonesia.

All `simple` and `uthmani` text are taken from Tanzil. The only modification done are removing basmalah
from beginning of each surah (except for surah Al-Fatiha) since basmalah is not part of ayah.

Unfortunately, their official download [link][tanzil-text-download] doesn't provide text with pause marks 
even if you set it otherwise. So, to download the complete text with pause marks and sajdah sign, we need
to use the unofficial API like this:

```
curl 'http://tanzil.net/tanzil/php/get-aya.php' \
	-H 'Referer: http://tanzil.net/' \
	--data-raw 'type=$TYPE&startAya=0&endAya=6236&version=1.5'
```

To use the above command, replace the `$TYPE` into one of the simple or Uthmani variant.

For `indonesia` text, it's taken from Quran Kemenag plugin for Microsoft Word which available in its
[official site][kemenag-quran]. If you want to extract the database by yourself, you can use install the
plugin then access the SQLite database in the installation location. If you are not using Windows, you can
extract the SQLite database by using [`innoextract`][innoextract] on the plugin installer.

## Translation

There are 148 translations which taken from several sources, mainly [Tanzil][tanzil-trans-download] and 
[QuranEnc][quranenc]. Besides those two sources, there is also translation that taken from 
[Quran Complex][quran-complex] which used for one of the Indonesian translation.

For the ones that taken from Tanzil, there are zero modification from the original sources except changing
format from CSV into JSON. For the ones that taken from QuranEnc, there some modifications to make sure all
footnotes have consistent numbering format, because currently they use several variations which differs
between each translation.

| No  	| ID                        	| Language    	| Title                                	| Translator                                                	| Source   	| Source Last Update  	|
|-----	|---------------------------	|-------------	|--------------------------------------	|-----------------------------------------------------------	|----------	|---------------------	|
| 1   	| sq-nahi-quranenc          	| Albanian    	| Albanian Translation                 	|                                                           	| QuranEnc 	| 2019-12-22          	|
| 2   	| sq-ahmeti-tanzil          	| Albanian    	| Sherif Ahmeti                        	| Sherif Ahmeti                                             	| Tanzil   	| 2010-09-14 04:38:47 	|
| 3   	| sq-nahi-tanzil            	| Albanian    	| Efendi Nahi                          	| Hasan Efendi Nahi                                         	| Tanzil   	| 2010-09-14 04:38:47 	|
| 4   	| ber-mensur-tanzil         	| Amazigh     	| At Mensur                            	| Ramdane At Mansour                                        	| Tanzil   	| 2012-08-01 04:00:02 	|
| 5   	| am-sadiq-quranenc         	| Amharic     	| Amharic translation                  	|                                                           	| QuranEnc 	| 2019-12-25          	|
| 6   	| am-sadiq-tanzil           	| Amharic     	| ሳዲቅ & ሳኒ ሐቢብ                         	| Muhammed Sadiq and Muhammed Sani Habib                    	| Tanzil   	| 2012-08-03 09:58:59 	|
| 7   	| ar-jalalayn-tanzil        	| Arabic      	| تفسير الجلالين                       	| Jalal ad-Din al-Mahalli and Jalal ad-Din as-Suyuti        	| Tanzil   	| 2010-12-01 07:16:03 	|
| 8   	| ar-muyassar-tanzil        	| Arabic      	| تفسير المیسر                         	| King Fahd Quran Complex                                   	| Tanzil   	| 2011-02-02 04:17:06 	|
| 9   	| as-rafeeq-quranenc        	| Assamese    	| Assamese translation                 	|                                                           	| QuranEnc 	| 2020-09-03          	|
| 10  	| az-mammadaliyev-tanzil    	| Azerbaijani 	| Məmmədəliyev & Bünyadov              	| Vasim Mammadaliyev and Ziya Bunyadov                      	| Tanzil   	| 2010-06-04 23:56:23 	|
| 11  	| az-musayev-tanzil         	| Azerbaijani 	| Musayev                              	| Alikhan Musayev                                           	| Tanzil   	| 2010-09-14 04:38:47 	|
| 12  	| bn-bengali-tanzil         	| Bengali     	| মুহিউদ্দীন খান                         	| Muhiuddin Khan                                            	| Tanzil   	| 2011-05-01 04:00:01 	|
| 13  	| bn-hoque-tanzil           	| Bengali     	| জহুরুল হক                              	| Zohurul Hoque                                             	| Tanzil   	| 2013-07-19 11:37:38 	|
| 14  	| bs-mihanovich-quranenc    	| Bosnian     	| Bosnian Translation                  	| Mihanovich                                                	| QuranEnc 	| 2019-12-21          	|
| 15  	| bs-rwwad-quranenc         	| Bosnian     	|  الترجمة البوسنية                    	| مركز رواد الترجمة                                         	| QuranEnc 	| 2020-03-03          	|
| 16  	| bs-korkut-tanzil          	| Bosnian     	| Korkut                               	| Besim Korkut                                              	| Tanzil   	| 2013-06-01 04:00:02 	|
| 17  	| bs-mlivo-tanzil           	| Bosnian     	| Mlivo                                	| Mustafa Mlivo                                             	| Tanzil   	| 2010-09-14 04:38:47 	|
| 18  	| bg-theophanov-tanzil      	| Bulgarian   	| Теофанов                             	| Tzvetan Theophanov                                        	| Tanzil   	| 2014-09-01 04:00:01 	|
| 19  	| zh-makin-quranenc         	| Chinese     	| Chinese Translation                  	|                                                           	| QuranEnc 	| 2017-07-13          	|
| 20  	| zh-jian-tanzil            	| Chinese     	| Ma Jian                              	| Ma Jian                                                   	| Tanzil   	| 2011-04-01 04:00:02 	|
| 21  	| zh-majian-tanzil          	| Chinese     	| Ma Jian (Traditional)                	| Ma Jian                                                   	| Tanzil   	| 2011-01-08 04:06:59 	|
| 22  	| cs-nykl-tanzil            	| Czech       	| Nykl                                 	| A. R. Nykl                                                	| Tanzil   	| 2010-09-14 04:38:47 	|
| 23  	| dv-divehi-tanzil          	| Divehi      	| ދިވެހި                                  	| Office of the President of Maldives                       	| Tanzil   	| 2011-03-02 21:21:37 	|
| 24  	| nl-keyzer-tanzil          	| Dutch       	| Keyzer                               	| Salomo Keyzer                                             	| Tanzil   	| 2010-06-04 23:56:23 	|
| 25  	| nl-leemhuis-tanzil        	| Dutch       	| Leemhuis                             	| Fred Leemhuis                                             	| Tanzil   	| 2012-08-05 17:14:53 	|
| 26  	| nl-siregar-tanzil         	| Dutch       	| Siregar                              	| Sofian S. Siregar                                         	| Tanzil   	| 2012-08-05 17:14:53 	|
| 27  	| en-hilali-quranenc        	| English     	| English Translation                  	| Hilali and Khan                                           	| QuranEnc 	| 2019-12-27          	|
| 28  	| en-rwwad-quranenc         	| English     	| English Translation                  	| Ruwwad Center - Amma Part                                 	| QuranEnc 	| 2021-01-18          	|
| 29  	| en-saheeh-quranenc        	| English     	| English Explanation                  	| Saheeh International                                      	| QuranEnc 	| 2019-12-24          	|
| 30  	| en-ahmedali-tanzil        	| English     	| Ahmed Ali                            	| Ahmed Ali                                                 	| Tanzil   	| 2010-09-14 04:38:47 	|
| 31  	| en-ahmedraza-tanzil       	| English     	| Ahmed Raza Khan                      	| Ahmed Raza Khan                                           	| Tanzil   	| 2011-08-01 04:00:02 	|
| 32  	| en-arberry-tanzil         	| English     	| Arberry                              	| A. J. Arberry                                             	| Tanzil   	| 2011-08-01 04:00:02 	|
| 33  	| en-daryabadi-tanzil       	| English     	| Daryabadi                            	| Abdul Majid Daryabadi                                     	| Tanzil   	| 2010-09-14 04:38:47 	|
| 34  	| en-hilali-tanzil          	| English     	| Hilali & Khan                        	| Muhammad Taqi-ud-Din al-Hilali and Muhammad Muhsin Khan   	| Tanzil   	| 2011-01-02 09:33:35 	|
| 35  	| en-itani-tanzil           	| English     	| Itani                                	| Talal Itani                                               	| Tanzil   	| 2013-07-19 11:37:38 	|
| 36  	| en-maududi-tanzil         	| English     	| Maududi                              	| Abul Ala Maududi                                          	| Tanzil   	| 2011-06-01 04:00:01 	|
| 37  	| en-mubarakpuri-tanzil     	| English     	| Mubarakpuri                          	| Safi-ur-Rahman al-Mubarakpuri                             	| Tanzil   	| 2015-06-20 21:51:34 	|
| 38  	| en-pickthall-tanzil       	| English     	| Pickthall                            	| Mohammed Marmaduke William Pickthall                      	| Tanzil   	| 2010-09-14 04:38:47 	|
| 39  	| en-qarai-tanzil           	| English     	| Qarai                                	| Ali Quli Qarai                                            	| Tanzil   	| 2013-07-16 18:55:21 	|
| 40  	| en-qaribullah-tanzil      	| English     	| Qaribullah & Darwish                 	| Hasan al-Fatih Qaribullah and Ahmad Darwish               	| Tanzil   	| 2010-06-04 23:56:23 	|
| 41  	| en-sahih-tanzil           	| English     	| Saheeh International                 	| Saheeh International                                      	| Tanzil   	| 2011-05-01 04:00:01 	|
| 42  	| en-sarwar-tanzil          	| English     	| Sarwar                               	| Muhammad Sarwar                                           	| Tanzil   	| 2012-08-01 04:00:02 	|
| 43  	| en-shakir-tanzil          	| English     	| Shakir                               	| Mohammad Habib Shakir                                     	| Tanzil   	| 2010-09-14 04:38:47 	|
| 44  	| en-transliteration-tanzil 	| English     	| Transliteration                      	| English Transliteration                                   	| Tanzil   	| 2010-09-14 04:38:47 	|
| 45  	| en-wahiduddin-tanzil      	| English     	| Wahiduddin Khan                      	| Wahiduddin Khan                                           	| Tanzil   	| 2011-08-25 13:06:08 	|
| 46  	| en-yusufali-tanzil        	| English     	| Yusuf Ali                            	| Abdullah Yusuf Ali                                        	| Tanzil   	| 2013-06-01 04:00:02 	|
| 47  	| fr-hameedullah-quranenc   	| French      	| French Translation                   	|                                                           	| QuranEnc 	| 2020-03-10          	|
| 48  	| fr-montada-quranenc       	| French      	| الترجمة الفرنسية                     	| مركز نور إنترناشونال                                      	| QuranEnc 	| 2018-10-11          	|
| 49  	| fr-hamidullah-tanzil      	| French      	| Hamidullah                           	| Muhammad Hamidullah                                       	| Tanzil   	| 2011-08-01 04:00:02 	|
| 50  	| de-bubenheim-quranenc     	| German      	| German Translation by Bubenheim      	|                                                           	| QuranEnc 	| 2021-01-07          	|
| 51  	| de-aburida-tanzil         	| German      	| Abu Rida                             	| Abu Rida Muhammad ibn Ahmad ibn Rassoul                   	| Tanzil   	| 2011-08-19 11:27:00 	|
| 52  	| de-bubenheim-tanzil       	| German      	| Bubenheim & Elyas                    	| A. S. F. Bubenheim and N. Elyas                           	| Tanzil   	| 2011-08-01 04:00:02 	|
| 53  	| de-khoury-tanzil          	| German      	| Khoury                               	| Adel Theodor Khoury                                       	| Tanzil   	| 2010-09-14 04:38:47 	|
| 54  	| de-zaidan-tanzil          	| German      	| Zaidan                               	| Amir Zaidan                                               	| Tanzil   	| 2010-06-04 23:56:23 	|
| 55  	| ha-gumi-quranenc          	| Hausa       	| Hausa language                       	|                                                           	| QuranEnc 	| 2021-01-07          	|
| 56  	| ha-gumi-tanzil            	| Hausa       	| Gumi                                 	| Abubakar Mahmoud Gumi                                     	| Tanzil   	| 2010-09-14 04:38:47 	|
| 57  	| hi-omari-quranenc         	| Hindi       	| Indian Translation                   	|                                                           	| QuranEnc 	| 2020-12-22          	|
| 58  	| hi-farooq-tanzil          	| Hindi       	| फ़ारूक़ ख़ान & अहमद                      	| Muhammad Farooq Khan and Muhammad Ahmed                   	| Tanzil   	| 2011-04-01 04:00:02 	|
| 59  	| hi-hindi-tanzil           	| Hindi       	| फ़ारूक़ ख़ान & नदवी                      	| Suhel Farooq Khan and Saifur Rahman Nadwi                 	| Tanzil   	| 2011-01-12 06:38:59 	|
| 60  	| Id-complex-1441           	| Indonesian  	| Indonesian Translation               	| King Fahd Quran Complex                                   	| KFQC     	| 2020-06-18          	|
| 61  	| id-affairs-quranenc       	| Indonesian  	| Indonesian Translation               	| Ministry of Islamic Affairs                               	| QuranEnc 	| 2018-08-13          	|
| 62  	| id-complex-quranenc       	| Indonesian  	| Indonesian Translation               	| King Fahd Quran Complex                                   	| QuranEnc 	| 2018-04-19          	|
| 63  	| id-sabiq-quranenc         	| Indonesian  	| Indonesian Translation               	| Sabeq Company                                             	| QuranEnc 	| 2019-12-19          	|
| 64  	| id-indonesian-tanzil      	| Indonesian  	| Bahasa Indonesia                     	| Indonesian Ministry of Religious Affairs                  	| Tanzil   	| 2010-06-04 23:56:23 	|
| 65  	| id-jalalayn-tanzil        	| Indonesian  	| Tafsir Jalalayn                      	| Jalal ad-Din al-Mahalli and Jalal ad-Din as-Suyuti        	| Tanzil   	| 2012-07-28 02:10:55 	|
| 66  	| id-muntakhab-tanzil       	| Indonesian  	| Quraish Shihab                       	| Muhammad Quraish Shihab et al.                            	| Tanzil   	| 2011-04-28 05:10:43 	|
| 67  	| it-piccardo-tanzil        	| Italian     	| Piccardo                             	| Hamza Roberto Piccardo                                    	| Tanzil   	| 2011-01-02 09:33:35 	|
| 68  	| ja-japanese-tanzil        	| Japanese    	| Japanese                             	| Unknown                                                   	| Tanzil   	| 2010-06-04 23:54:20 	|
| 69  	| km-cambodia-quranenc      	| Khmer       	| Khmer translation                    	|                                                           	| QuranEnc 	| 2017-02-18          	|
| 70  	| ko-korean-tanzil          	| Korean      	| Korean                               	| Unknown                                                   	| Tanzil   	| 2011-08-01 04:00:02 	|
| 71  	| ku-bamoki-quranenc        	| Kurdish     	| Kurdish Translation                  	|                                                           	| QuranEnc 	| 2019-12-28          	|
| 72  	| ms-basmeih-tanzil         	| Malay       	| Basmeih                              	| Abdullah Muhammad Basmeih                                 	| Tanzil   	| 2012-10-01 04:00:02 	|
| 73  	| ml-abdulhameed-tanzil     	| Malayalam   	| അബ്ദുല്‍ ഹമീദ് & പറപ്പൂര്‍                    	| Cheriyamundam Abdul Hameed and Kunhi Mohammed Parappoor   	| Tanzil   	| 2012-05-01 04:00:01 	|
| 74  	| ml-karakunnu-tanzil       	| Malayalam   	| കാരകുന്ന് & എളയാവൂര്                      	| Muhammad Karakunnu and Vanidas Elayavoor                  	| Tanzil   	| 2012-05-01 04:00:01 	|
| 75  	| no-berg-tanzil            	| Norwegian   	| Einar Berg                           	| Einar Berg                                                	| Tanzil   	| 2010-06-04 23:56:23 	|
| 76  	| om-ababor-quranenc        	| Oromo       	| Oromo translation                    	|                                                           	| QuranEnc 	| 2017-03-19          	|
| 77  	| ps-abdulwali-tanzil       	| Pashto      	| عبدالولي                             	| Abdulwali Khan                                            	| Tanzil   	| 2016-07-01 04:00:01 	|
| 78  	| fa-ih-quranenc            	| Persian     	| Persian Translation                  	| Islamhouse                                                	| QuranEnc 	| 2020-05-10          	|
| 79  	| fa-tagi-quranenc          	| Persian     	| Persian translation                  	| Husein Tagy                                               	| QuranEnc 	| 2018-08-14          	|
| 80  	| fa-ansarian-tanzil        	| Persian     	| انصاریان                             	| Hussain Ansarian                                          	| Tanzil   	| 2011-08-01 04:00:02 	|
| 81  	| fa-ayati-tanzil           	| Persian     	| آیتی                                 	| AbdolMohammad Ayati                                       	| Tanzil   	| 2012-09-01 04:00:02 	|
| 82  	| fa-bahrampour-tanzil      	| Persian     	| بهرام پور                            	| Abolfazl Bahrampour                                       	| Tanzil   	| 2011-04-14 10:09:03 	|
| 83  	| fa-fooladvand-tanzil      	| Persian     	| فولادوند                             	| Mohammad Mahdi Fooladvand                                 	| Tanzil   	| 2011-07-01 04:00:01 	|
| 84  	| fa-gharaati-tanzil        	| Persian     	| قرائتی                               	| Mohsen Gharaati                                           	| Tanzil   	| 2017-12-31 12:12:07 	|
| 85  	| fa-ghomshei-tanzil        	| Persian     	| الهی قمشه‌ای                          	| Mahdi Elahi Ghomshei                                      	| Tanzil   	| 2011-08-01 04:00:02 	|
| 86  	| fa-khorramdel-tanzil      	| Persian     	| خرمدل                                	| Mostafa Khorramdel                                        	| Tanzil   	| 2011-03-19 20:35:53 	|
| 87  	| fa-khorramshahi-tanzil    	| Persian     	| خرمشاهی                              	| Baha'oddin Khorramshahi                                   	| Tanzil   	| 2012-07-27 23:20:01 	|
| 88  	| fa-makarem-tanzil         	| Persian     	| مکارم شیرازی                         	| Naser Makarem Shirazi                                     	| Tanzil   	| 2014-02-01 05:00:02 	|
| 89  	| fa-moezzi-tanzil          	| Persian     	| معزی                                 	| Mohammad Kazem Moezzi                                     	| Tanzil   	| 2010-09-14 04:38:47 	|
| 90  	| fa-mojtabavi-tanzil       	| Persian     	| مجتبوی                               	| Sayyed Jalaloddin Mojtabavi                               	| Tanzil   	| 2012-05-01 04:00:01 	|
| 91  	| fa-sadeqi-tanzil          	| Persian     	| صادقی تهرانی                         	| Mohammad Sadeqi Tehrani                                   	| Tanzil   	| 2011-09-01 04:00:01 	|
| 92  	| pl-bielawskiego-tanzil    	| Polish      	| Bielawskiego                         	| Józefa Bielawskiego                                       	| Tanzil   	| 2010-09-14 04:38:47 	|
| 93  	| pt-nasr-quranenc          	| Portuguese  	| Portuguese Translation               	|                                                           	| QuranEnc 	| 2020-09-22          	|
| 94  	| pt-elhayek-tanzil         	| Portuguese  	| El-Hayek                             	| Samir El-Hayek                                            	| Tanzil   	| 2010-06-04 23:56:23 	|
| 95  	| ro-grigore-tanzil         	| Romanian    	| Grigore                              	| George Grigore                                            	| Tanzil   	| 2010-09-14 04:38:47 	|
| 96  	| ru-abuadel-tanzil         	| Russian     	| Абу Адель                            	| Abu Adel                                                  	| Tanzil   	| 2010-09-16 02:26:15 	|
| 97  	| ru-krachkovsky-tanzil     	| Russian     	| Крачковский                          	| Ignaty Yulianovich Krachkovsky                            	| Tanzil   	| 2010-09-20 05:33:44 	|
| 98  	| ru-kuliev-alsaadi-tanzil  	| Russian     	| Кулиев + ас-Саади                    	| Elmir Kuliev (with Abd ar-Rahman as-Saadi's commentaries) 	| Tanzil   	| 2010-10-04 06:30:28 	|
| 99  	| ru-kuliev-tanzil          	| Russian     	| Кулиев                               	| Elmir Kuliev                                              	| Tanzil   	| 2011-06-01 04:00:01 	|
| 100 	| ru-muntahab-tanzil        	| Russian     	| Аль-Мунтахаб                         	| Ministry of Awqaf, Egypt                                  	| Tanzil   	| 2011-06-01 04:00:01 	|
| 101 	| ru-osmanov-tanzil         	| Russian     	| Османов                              	| Magomed-Nuri Osmanovich Osmanov                           	| Tanzil   	| 2010-09-14 04:38:47 	|
| 102 	| ru-porokhova-tanzil       	| Russian     	| Порохова                             	| V. Porokhova                                              	| Tanzil   	| 2010-09-14 04:38:47 	|
| 103 	| ru-sablukov-tanzil        	| Russian     	| Саблуков                             	| Gordy Semyonovich Sablukov                                	| Tanzil   	| 2010-10-04 06:30:28 	|
| 104 	| sd-amroti-tanzil          	| Sindhi      	| امروٽي                               	| Taj Mehmood Amroti                                        	| Tanzil   	| 2012-05-01 04:00:01 	|
| 105 	| si-mahir-quranenc         	| Sinhalese   	| الترجمة السنهالية                    	|                                                           	| QuranEnc 	| 2020-06-26          	|
| 106 	| so-abduh-tanzil           	| Somali      	| Abduh                                	| Mahmud Muhammad Abduh                                     	| Tanzil   	| 2010-09-14 04:38:47 	|
| 107 	| es-garcia-quranenc        	| Spanish     	| Spanish Translation                  	|                                                           	| QuranEnc 	| 2016-12-28          	|
| 108 	| es-montada-eu-quranenc    	| Spanish     	| الترجمة الإسبانية                    	| مركز نور إنترناشونال                                      	| QuranEnc 	| 2018-10-09          	|
| 109 	| es-montada-latin-quranenc 	| Spanish     	| الترجمة الإسبانية (أمريكا اللاتينية) 	| مركز نور إنترناشونال                                      	| QuranEnc 	| 2018-10-09          	|
| 110 	| es-bornez-tanzil          	| Spanish     	| Bornez                               	| Raúl González Bórnez                                      	| Tanzil   	| 2014-07-15 18:04:37 	|
| 111 	| es-cortes-tanzil          	| Spanish     	| Cortes                               	| Julio Cortes                                              	| Tanzil   	| 2014-09-01 04:00:01 	|
| 112 	| es-garcia-tanzil          	| Spanish     	| Garcia                               	| Muhammad Isa García                                       	| Tanzil   	| 2014-07-15 18:04:37 	|
| 113 	| sw-barwani-tanzil         	| Swahili     	| Al-Barwani                           	| Ali Muhsin Al-Barwani                                     	| Tanzil   	| 2010-09-14 04:38:47 	|
| 114 	| sv-bernstrom-tanzil       	| Swedish     	| Bernström                            	| Knut Bernström                                            	| Tanzil   	| 2012-08-01 04:00:02 	|
| 115 	| tl-rwwad-quranenc         	| Tagalog     	| الترجمة الفلبينية (تجالوج)           	|                                                           	| QuranEnc 	| 2020-06-29          	|
| 116 	| tg-arifi-quranenc         	| Tajik       	| الترجمة الطاجيكية                    	| عارفي                                                     	| QuranEnc 	| 2018-09-29          	|
| 117 	| tg-ayati-tanzil           	| Tajik       	| Оятӣ                                 	| AbdolMohammad Ayati                                       	| Tanzil   	| 2010-09-14 04:38:47 	|
| 118 	| ta-baqavi-quranenc        	| Tamil       	| الترجمة التاميلية                    	| عبد الحميد باقوي                                          	| QuranEnc 	| 2021-01-07          	|
| 119 	| ta-tamil-tanzil           	| Tamil       	| ஜான் டிரஸ்ட்                            	| Jan Turst Foundation                                      	| Tanzil   	| 2010-09-14 04:38:47 	|
| 120 	| tt-nugman-tanzil          	| Tatar       	| Yakub Ibn Nugman                     	| Yakub Ibn Nugman                                          	| Tanzil   	| 2010-09-14 04:38:47 	|
| 121 	| th-thai-tanzil            	| Thai        	| ภาษาไทย                              	| King Fahad Quran Complex                                  	| Tanzil   	| 2011-11-01 04:00:02 	|
| 122 	| tr-rwwad-quranenc         	| Turkish     	| الترجمة التركية                      	| مركز رواد الترجمة                                         	| QuranEnc 	| 2018-10-16          	|
| 123 	| tr-shaban-quranenc        	| Turkish     	| Turkish translation                  	| Shaaban Britsh                                            	| QuranEnc 	| 2019-12-26          	|
| 124 	| tr-shahin-quranenc        	| Turkish     	| الترجمة التركية                      	| د. علي أوزك وآخرون                                        	| QuranEnc 	| 2017-05-23          	|
| 125 	| tr-ates-tanzil            	| Turkish     	| Süleyman Ateş                        	| Suleyman Ates                                             	| Tanzil   	| 2010-06-04 23:56:23 	|
| 126 	| tr-bulac-tanzil           	| Turkish     	| Alİ Bulaç                            	| Alİ Bulaç                                                 	| Tanzil   	| 2010-09-14 04:38:47 	|
| 127 	| tr-diyanet-tanzil         	| Turkish     	| Diyanet İşleri                       	| Diyanet Isleri                                            	| Tanzil   	| 2012-01-01 05:00:01 	|
| 128 	| tr-golpinarli-tanzil      	| Turkish     	| Abdulbakî Gölpınarlı                 	| Abdulbaki Golpinarli                                      	| Tanzil   	| 2010-06-04 23:54:20 	|
| 129 	| tr-ozturk-tanzil          	| Turkish     	| Öztürk                               	| Yasar Nuri Ozturk                                         	| Tanzil   	| 2010-06-04 23:56:23 	|
| 130 	| tr-transliteration-tanzil 	| Turkish     	| Çeviriyazı                           	| Muhammet Abay                                             	| Tanzil   	| 2010-09-16 02:26:15 	|
| 131 	| tr-vakfi-tanzil           	| Turkish     	| Diyanet Vakfı                        	| Diyanet Vakfi                                             	| Tanzil   	| 2010-06-05 00:00:03 	|
| 132 	| tr-yazir-tanzil           	| Turkish     	| Elmalılı Hamdi Yazır                 	| Elmalili Hamdi Yazir                                      	| Tanzil   	| 2010-06-05 00:00:03 	|
| 133 	| tr-yildirim-tanzil        	| Turkish     	| Suat Yıldırım                        	| Suat Yildirim                                             	| Tanzil   	| 2010-06-05 00:00:03 	|
| 134 	| tr-yuksel-tanzil          	| Turkish     	| Edip Yüksel                          	| Edip Yüksel                                               	| Tanzil   	| 2010-06-04 23:56:23 	|
| 135 	| ur-junagarhi-quranenc     	| Urdu        	| Urdu Translation                     	|                                                           	| QuranEnc 	| 2021-01-19          	|
| 136 	| ur-ahmedali-tanzil        	| Urdu        	| احمد علی                             	| Ahmed Ali                                                 	| Tanzil   	| 2010-09-14 04:38:47 	|
| 137 	| ur-jalandhry-tanzil       	| Urdu        	| جالندہری                             	| Fateh Muhammad Jalandhry                                  	| Tanzil   	| 2011-01-02 09:33:35 	|
| 138 	| ur-jawadi-tanzil          	| Urdu        	| علامہ جوادی                          	| Syed Zeeshan Haider Jawadi                                	| Tanzil   	| 2011-01-02 09:33:35 	|
| 139 	| ur-junagarhi-tanzil       	| Urdu        	| محمد جوناگڑھی                        	| Muhammad Junagarhi                                        	| Tanzil   	| 2011-04-25 23:56:51 	|
| 140 	| ur-kanzuliman-tanzil      	| Urdu        	| احمد رضا خان                         	| Ahmed Raza Khan                                           	| Tanzil   	| 2011-04-01 04:00:02 	|
| 141 	| ur-maududi-tanzil         	| Urdu        	| ابوالاعلی مودودی                     	| Abul A'ala Maududi                                        	| Tanzil   	| 2010-12-01 07:16:03 	|
| 142 	| ur-najafi-tanzil          	| Urdu        	| محمد حسین نجفی                       	| Muhammad Hussain Najafi                                   	| Tanzil   	| 2011-08-19 11:27:00 	|
| 143 	| ur-qadri-tanzil           	| Urdu        	| طاہر القادری                         	| Tahir ul Qadri                                            	| Tanzil   	| 2010-09-14 04:38:47 	|
| 144 	| ug-saleh-quranenc         	| Uyghur      	| Uyghur translation                   	|                                                           	| QuranEnc 	| 2018-02-20          	|
| 145 	| ug-saleh-tanzil           	| Uyghur      	| محمد صالح                            	| Muhammad Saleh                                            	| Tanzil   	| 2010-06-05 00:00:03 	|
| 146 	| uz-mansour-quranenc       	| Uzbek       	| Uzbek translation                    	| Alauddin Mansour                                          	| QuranEnc 	| 2017-03-25          	|
| 147 	| uz-sodik-tanzil           	| Uzbek       	| Мухаммад Содик                       	| Muhammad Sodik Muhammad Yusuf                             	| Tanzil   	| 2011-08-01 04:00:02 	|
| 148 	| yo-mikail-quranenc        	| Yoruba      	| Yoruba translation                   	|                                                           	| QuranEnc 	| 2020-10-20          	|

There are several translations from Tanzil and QuranEnc that excluded from this repository because they
are still incomplete. From Tanzil, these are the excluded translations:

- Feti Mehdiu (Albanian), missing translation for Al-Anbiya (21):56 and Al-Mursilat (77):14
- Hrbek (Czech), missing for Ar-Rum (30):18, Al-Waqi'a (56):2 and Al-Waqi'a (56):13
- Burhan Muhammad-Amin (Kurdish), missing for Al-Kautsar (108):3

From QuranEnc, Kinyarwanda translation is excluded because it's missing translation for:

- Fussilat (41):27
- Ash-Shura (42):18
- Az-Zukhruf (43):14
- Az-Zukhruf (43):68
- Al-Fath (48):2
- Adh-Dhariyat (51):6
- Ar-Rahman (55):14

[tanzil-meta]: http://tanzil.net/docs/quran_metadata
[tanzil-text-download]: http://tanzil.net/pub/download/download.htm
[tanzil-trans-download]: http://tanzil.net/trans/
[quranenc]: https://quranenc.com/en/home
[quran-complex]: https://qurancomplex.gov.sa/kfgqpc-quran-translate/
[kemenag-quran]: https://lajnah.kemenag.go.id/unduhan/category/1-qkiw
[innoextract]: https://constexpr.org/innoextract/