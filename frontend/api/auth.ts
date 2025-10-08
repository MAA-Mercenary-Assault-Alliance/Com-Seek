// src/api/auth.ts
import { api } from "./client";

export type UserType = "student" | "company";

export async function login(email: string, password: string) {
  const { data } = await api.post("/auth/login", { email, password });
  return data; // { message, email }
}

type StudentProfile = {
  first_name: string;
  last_name: string;
  is_alum: boolean;
};

type CompanyProfile = {
  name: string;
  location: string;
  contact_email: string;
  contact_number: string;
};

export async function registerStudent(payload: {
  email: string;
  password: string;
  student: StudentProfile;
}) {
  const body = {
    email: payload.email,
    password: payload.password,
    user_type: "student",
    student_profile: payload.student,
  };
  const { data } = await api.post("/auth/register", body);
  return data; // { data: user }
}

export async function registerCompany(payload: {
  email: string;
  password: string;
  company: CompanyProfile;
}) {
  const body = {
    email: payload.email,
    password: payload.password,
    user_type: "company",
    company_profile: payload.company,
  };
  const { data } = await api.post("/auth/register", body);
  return data; // { data: user }
}

// quick auth-check by hitting a protected endpoint
export async function getMyStudentProfile() {
  const { data } = await api.get("/student/");
  return data;
}

export async function getMyCompanyProfile() {
  const { data } = await api.get("/company/");
  return data;
}
