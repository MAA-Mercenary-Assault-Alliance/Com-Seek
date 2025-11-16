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
        profile_image_id?: string | null;
        cover_image_id?: string | null;
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

export interface CompanyTemplate {
    UserID: number;
    Name: string;
    Website: string;
    ContactEmail: string;
    ContactNumber: string;
    Location: string;
    Description: string;
    Approved: boolean;
    profile_image_id?: string | null;
    cover_image_id?: string | null;
    Jobs: any[]; // replace 'any' with Job[] if you define a Job interface later
    CreatedAt?: string; // optional if provided by backend
}



export interface StudentTemplate {
    UserID: number;
    User: {
        ID: number;
        CreatedAt: string;
        UpdatedAt: string;
        DeletedAt: string | null;
        Email: string;
        Password: string;
    };
    CreatedAt: string;
    FirstName: string;
    LastName: string;
    Description: string;
    IsAlum: boolean;
    Approved: boolean;
    GitHub: string;
    LinkedIn: string;
    Facebook: string;
    Instagram: string;
    Twitter: string;
    profile_image_id?: string | null;
    cover_image_id?: string | null;
    JobApplication: any; // or a specific type if you know its structure
}

