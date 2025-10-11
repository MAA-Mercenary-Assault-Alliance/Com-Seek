const jobMarkdown = `
The Account Manager is responsible for building and maintaining strong client relationships while driving sales growth. This role involves identifying customer needs, presenting tailored solutions, and ensuring long-term satisfaction. The Account Manager acts as the main point of contact between the company and its clients, balancing technical knowledge with strong communication skills to deliver value.

**Key Responsibilities:**

- Develop and manage client accounts to achieve sales targets.
- Identify opportunities for upselling and cross-selling.
- Collaborate with technical teams to provide tailored solutions.
- Maintain regular communication to ensure customer satisfaction.
- Prepare reports, proposals, and sales forecasts.

**Skills & Qualifications:**

- Strong interpersonal and negotiation skills.
- Technical knowledge relevant to the industry.
- Ability to manage multiple accounts and priorities.
- Excellent problem-solving and communication abilities.
`

export const TE_Info = {
    id: "1",
    name: "ACCOUNT MANAGER (Sales Engineer)",
    company: "TE Connectivity",
    location: "กรุงเทพมหานคร",
    salary_range: "฿12000-฿15000 per month", // MAKE THIS FLOOR AND CEILING
    type: "full-time",
    experience: "1-2 years",
    desc: jobMarkdown,
}

export interface JobTemplate {
    ID: number;
    Title: string;
    CompanyID: number;
    Company: {
        UserID: number;
        User: {
            ID: number;
            CreatedAt: string;
            UpdatedAt: string;
            DeletedAt: string | null;
            email: string;
            password: string;
        };
        Name: string;
        Website: string;
        ContactEmail: string;
        ContactNumber: string;
        Location: string;
        Description: string;
        Approved: boolean;
        Jobs: JobTemplate[] | null;
    };
    Location: string;
    JobType: string;
    EmploymentStatus: string;
    MinSalary: number;
    MaxSalary: number;
    MinExperience: number;
    MaxExperience: number;
    Description: string;
    CreatedAt: string;
    UpdatedAt: string;
    Approved: boolean;
    Visibility: boolean;
    CheckNeeded: boolean;
}



export const TE_Info_Company = {
    name: "TE Connectivity",
    website: "some url",
    location: "กรุงเทพมหานคร",
    description: jobMarkdown,
    createdAt: "16/07/2025",
    approved: true,
    number: "111111111",
    email: "TeConnecc@gmail.com",
}

export interface CompanyTemplate {
    name: string;
    website: string;
    location: string;
    description: string;
    createdAt: string;
    approved: boolean;
    number: string;
    email: string;
}

export const Student_Info = {
    id: "11",
    user: "what",
    firstName: "Nagi",
    lastName: "Seishiro",
    profile_pic: "/nagi",
    description: "I can even destroy blue lock",
    approved: true,
}

export interface StudentTemplate {
    id: string;
    user: string;
    firstName: string;
    lastName: string;
    profile_pic: string;
    description: string;
    approved: boolean;
}
