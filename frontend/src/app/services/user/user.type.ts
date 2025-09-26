export interface UserResponse {
  id: number;
  wcaId: string | null;
  name: string;
  email: string;
  gender: string;
  bio: string;
  countryISO: string;
  avatarURL: string;
  role: string;
  createdAt: Date;
}

export interface UpdateUserBioRequestBody {
  bio: string;
}
