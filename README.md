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

There are [152](translation) translations which taken from several sources, mainly [Tanzil][tanzil-trans-download] 
and [QuranEnc][quranenc]. Besides those two sources, there is also translation that taken from
[Quran Complex][quran-complex] which used for one of the Indonesian translation.

For the ones that taken from Tanzil, there are zero modification from the original sources except changing
format from CSV into JSON. For the ones that taken from QuranEnc, there some modifications to make sure all
footnotes have consistent numbering format, because currently they use several variations which differs
between each translation.

There are several translations from Tanzil and QuranEnc that incomplete, but still included in this repository
because they are still useful and could be corrected later. From Tanzil, these are the incomplete translations:

- `sq-mehdiu` by Feti Mehdiu (Albanian), missing translation for Al-Anbiya (21):56 and Al-Mursilat (77):14
- `cs-hrbek` by Preklad I. Hrbek (Czech), missing Ar-Rum (30):18, Al-Waqi'a (56):2 and Al-Waqi'a (56):13
- `ku-asan` by Burhan Muhammad-Amin (Kurdish), missing Al-Kautsar (108):3

From QuranEnc, Kinyarwanda translation (`rw-assoc`) is incomplete because it's missing translation for:

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