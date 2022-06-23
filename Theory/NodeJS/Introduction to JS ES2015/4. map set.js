"use strict";

const profiles = new Map();
profiles.set("twitter", "@adalovelace");
profiles.set("facebook", "adalovelace");
profiles.set("googleplus", "ada");
profiles.size; // 3
profiles.has("twitter"); // true
profiles.get("twitter"); // "@adalovelace"
profiles.has("youtube"); // false
profiles.delete("facebook");
profiles.has("facebook"); // false
profiles.get("facebook"); // undefined
for (const entry of profiles) {
  console.log(entry);
}

// second

const tests = new Map();
tests.set(() => 2 + 2, 4);
tests.set(() => 2 * 2, 4);
tests.set(() => 2 / 2, 1);
tests.set(() => 2 + 2, 5);
for (const entry of tests) {
  console.log(entry[0]() === entry[1] ? "PASS" : "FAIL");
}

//third

const s = new Set([0, 1, 2, 3]);
s.add(3); // не будет добавлено
s.size; // 4
s.delete(0);
s.has(0); // false
for (const entry of s) {
  console.log(entry);
}
