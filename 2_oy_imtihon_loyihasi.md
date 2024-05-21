# Online Bookstore API Examination (Onlayn kitob do'koni API imtihoni)

## Project Overview (Loyihaga umumiy nuqtai)

Sizga `Go` yordamida onlayn kitob do'koni uchun `RESTful API` ishlab chiqish vazifasi yuklangan. Kitoblar (books) va mualliflarni (authors)boshqarish uchun `endpoint` larni taqdim etishingiz kerak. `API` `PostgreSQL` ma'lumotlar bazasi bilan ishlashga mo'ljallangan bo'lishi kerak.

## Mal'umotlar bazasigadi sxemasi

## Database name: bookstore_n9
<!-- ## Postgres user: student_n9, postgres password: student_n9 -->
## Add a mock data (at least 10 values for each table)
## Authors Table Fields (Mualliflar jadvali ustunlari)

1. **author_id (Primary Key):**
   - Har bir muallif uchun noyob identifikator.
   - Type: Integer or UUID (Universally Unique Identifier).

2. **name:**
   - Muallifning ismi.
   - Type: String.

3. **birth_date:**
   - Muallifning tug'ilgan sanasi.
   - Type: Date.

4. **biography:**
   - Muallifning qisqacha tarjimai holi yoki tavsifi.
   - Type: Text.

5. **created_at:**
   - Muallif yozuvi yaratilgan vaqt tamg'asi.
   - Type: Timestamp or Datetime.

6. **updated_at:**
   - Muallif yozuvi oxirgi marta yangilangan vaqt belgisi.
   - Type: Timestamp or Datetime.

## Books Table Fields (Kitoblar jadvali ustunlari)

1. **book_id (Primary Key):**
   - Har bir kitob uchun noyob identifikator.
   - Type: Integer or UUID.

2. **title:**
   - Kitobning nomi.
   - Type: String.

3. **author_id (Foreign Key):**
   - Kitobni yozgan muallifga havola.
   - Type: Integer or UUID.

4. **publication_date:**
   - Kitobning nashr etilgan sanasi.
   - Type: Date.

5. **isbn:**
   - Kitobning xalqaro standart kitob raqami (ISBN).
   - Type: String.

6. **description:**
   - Kitobning qisqacha tavsifi yoki xulosasi.
   - Type: Text or String.

7. **created_at:**
   - Kitob yozuvi yaratilgan vaqt tamg'asi.
   - Type: Timestamp or Datetime.

8. **updated_at:**
   - Kitob yozuvi oxirgi marta yangilangan vaqt belgisi.
   - Type: Timestamp or Datetime.

---
### Project Requirements (Loyiha talablari)

1. **Migrations (5 points):**
   - Sxema o'zgarishlarini boshqarish uchun ma'lumotlar bazasi migratsiyasini amalga oshiring.

2. **Books Management (20 points):**
   - Kitoblar uchun `CRUD` operatsiyalari (`create`, `read (all, by id)`, `update`, `delete`). (**15 points**)
   - Search and filter functionality. (Qidiruv va filtrlash) (**5 points**)

3. **Authors Management (20 points):**
   - mualliflar uchun `CRUD` operatsiyalari. (`create`, `read (all, by id)`, `update`, `delete`). (**15 points**)
   - Search and filter functionality. (Qidiruv va filtrlash) (**5 points**)

4. **Relationship between Books and Authors (10 points):**
   - Ma'lumotlar bazasi sxemasida kitoblar va mualliflar o'rtasidagi munosabatlarni o'rnatish.

5. **Configuration (5 points):**
   - Ilova uchun konfiguratsiya boshqaruvini amalga oshiring.

6. **Makefile (5 points):**
   - Ilovani sinab ko'rish va ishga tushirish uchun `Makefile` yarating.

7. **Bonus points:**
   - Code layout (**5 points**)
   - Request logging, error handling, content-type tekshirish va boshqalar kabi `middleware` komponentlarini amalga oshiring. (**10 points**)
   - So'rovlar ishlashini optimallashtirish uchun ma'lumotlar bazasi sxemasida indekslardan foydalaning. (**10 points**)
   - To'g'ridan-to'g'ri "`net/http`" kutubxonasidan uchinchi tomon vosita dasturlari yoki frameworklar tayanmasdan foydalanadigan vaziyat (**10 points**)



## Assessment Criteria (Baholash mezonlari)
Loyihangiz quyidagi mezonlar asosida baholanadi:

- Loyiha talablari va ko'rsatmalariga rioya qilish.
- `API` funksionalligining to'liqligi va to'g'riligi.
- `RESTful API `tamoyillari va eng yaxshi amaliyotlarga muvofiqligi.
- `API`ni tekshirish uchun ma'lumotlarni tayyorlang
- Ma'lumotlar bazasi migratsiyasini, konfiguratsiyani va `Makefile`-ni amalga oshirish.
- Kod sifati, tashkiliyligi va o'qilishi



