export async function fetchStudentProfile() {
  // simulate network delay
  await new Promise(r => setTimeout(r, 500));

  return {
    firstName: "Reed",
    lastName: "Richard",
    username: "Fantastic-Username",
    email: "mybigthing@gmail.com",
    joinDate: "2019-04-13",
    bio: "This is going to be fantastic!",
    avatar: "/src/assets/avatar.png",
    cover: "/src/assets/cover.jpeg",
    socials: {
      facebook: "https://facebook.com",
      twitter: "https://twitter.com",
      instagram: "https://instagram.com",
      github: "https://github.com"
    }
  };
}
