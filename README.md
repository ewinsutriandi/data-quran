# Al-Quran Database

This repository contains all data that is related with Al-Quran, from the text, translation, explanation
(tafseer), metadata and font. It's collected from several sources and stored as JSON (except for font,
obviously) to make it easy to use in various programming language.

## Metadata

Metadata is taken from [Tanzil][tanzil-meta] with minor modification. There are several metadatas available:

- Surah, the chapter of the Quran. There are 114 surahs in the Quran, each divided into ayahs (verses)
- Juz, the division of Quran into 30 parts.
- Hizb, the division of Quran where one juz has two hizb. The hizb itself sometimes is separated into four 
  quarters, so in Quran there are 240 hizb quarters.
- Ruku basically is a paragraph of the Quran. Used to denote a group of thematically related verses in the
  Quran, so that the reciters could identify when to make ruku in Salah without breaking an ongoing topic
  in the Quranic text.
- Page is division of Quran page based on Mushaf Madinah.

[tanzil-meta]: http://tanzil.net/docs/quran_metadata