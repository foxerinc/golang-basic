# Data Type Exercise — Answer Notes

This document summarizes reasonable data type choices for each case and the assumptions behind them.  
I’ll use **generic programming types**: integer, floating‑point, string, boolean.

---

## 1. Person’s age

**Suggested type:** Integer (e.g. `int`)

**Reasoning:**

- Age in years is a whole number (20, 35, 72).
- You do not need fractional values like 20.5 years in most apps.
- Integer is simple, efficient, and clear for counting years.

**Assumption:**  
You only care about full years, not months or days.

---

## 2. Number of pages in a book

**Suggested type:** Integer (e.g. `int`)

**Reasoning:**

- Page counts are whole numbers (150 pages, 432 pages).
- Integer matches a count of discrete items.

**Assumption:**  
Books will not have so many pages that you overflow a normal integer.

---

## 3. Price of a meal at a restaurant

**Suggested type:** Decimal‑like numeric or fixed‑point (e.g. `decimal`, or integer of smallest currency unit like cents)

**Reasoning:**

- Prices often use 2 decimal places (e.g. 12.50).
- Binary floating‑point types (`float`) can introduce rounding errors for money.
- Using a decimal type, or storing in smallest unit (cents) as integer, avoids precision issues.

**Assumption:**  
You care about accurate financial calculations, not just approximate display.

---

## 4. Total population of a city

**Suggested type:** Large integer (e.g. 64‑bit integer, `int64` / `long`)

**Reasoning:**

- Population is a whole number.
- Big cities can have millions of people; some datasets can aggregate even more.
- A wide integer type gives enough range for safety.

**Assumption:**  
You might track very large populations, so you choose a type with plenty of headroom.

---

## 5. Grade Point Average (GPA)

**Suggested type:** Floating‑point or decimal (e.g. `float`, `double`, or `decimal`)

**Reasoning:**

- GPA commonly has decimals (e.g. 3.25, 3.87).
- You need fractional precision.
- A float is usually enough; a decimal type is safer if you need very consistent rounding.

**Assumption:**  
You accept standard floating‑point behavior or a decimal type is available if you need stricter precision.

---

## 6. Email address

**Suggested type:** String

**Reasoning:**

- Email addresses are a mix of characters: letters, digits, `@`, `.`, etc.
- They are identifiers, not values for arithmetic.
- String is the natural representation.

**Assumption:**  
Validation is done separately (regex, library), not by the data type itself.

---

## 7. User accepted terms and conditions (flag)

**Suggested type:** Boolean (`true` / `false`)

**Reasoning:**

- The state is binary: accepted or not accepted.
- Boolean directly communicates this intent.
- More readable than `0/1` integers or `"yes"/"no"` strings.

**Assumption:**  
You only need two states. If you later need “not answered yet”, you might wrap the boolean in a struct or allow a nullable boolean.

---

## 8. Distance from the Earth to the Moon

**Suggested type:** Floating‑point (e.g. `double`, `float64`)

**Reasoning:**

- Distance is a physical measurement and can vary slightly.
- You might store it with units like kilometers or meters, possibly with fractions.
- A float gives range and fractional precision.

**Assumption:**  
You need approximate scientific/engineering accuracy, not exact integer counting of meters.

---

## 9. ISBN number of a book

**Suggested type:** String

**Reasoning:**

- ISBNs can contain leading zeros and hyphens (e.g. `978-3-16-148410-0`).
- You never perform arithmetic on an ISBN.
- Treating it as a number can strip leading zeros and cause confusion.

**Assumption:**  
You only compare, format, or validate ISBNs, not compute with them.

---

## 10. Hexadecimal value of a color

**Suggested type:** String

**Reasoning:**

- Hex color codes are written as text (e.g. `#FF5733`, `#000000`).
- The leading `#` and hex digits are part of the representation.
- A string keeps the full format as‑is.

**Alternative:**  
You could parse into separate integer channels (R, G, B, A) internally, but the hex itself is best stored as string.

---

## 11. Rating of a movie (1–5 stars)

**Suggested type:** Small integer (e.g. `int`)

**Reasoning:**

- Ratings are usually discrete values (1, 2, 3, 4, 5).
- Integer matches this scale and is easy to validate (`1 <= rating <= 5`).

**Assumption:**  
Users pick whole stars only. If you later support half‑stars (e.g. 4.5), you might switch to a decimal or store as “half‑star units” in an integer.

---

## 12. File size of an image in bytes

**Suggested type:** Unsigned or large integer (e.g. `uint64` / `long`)

**Reasoning:**

- File size is a count of bytes, so it is always non‑negative.
- Large files (videos, high‑res images) can be many bytes, so a wider type is safer.
- Integer also makes it easy to convert to KB/MB/GB when displaying.

**Assumption:**  
You want to support big files without overflow and you do not need negative values.

---

## Short Summary

- **Counts / whole numbers:** age, pages, population, rating, file size → **integer**.
- **Money / measurements / averages:** price, GPA, distance → **decimal / float**, with preference for **decimal or integer of smallest unit** for money.
- **Identifiers / text:** email, ISBN, hex color → **string**.
- **Flags / yes‑no:** accepted terms → **boolean**.

These choices keep the model simple and close to how the data is used in real applications.