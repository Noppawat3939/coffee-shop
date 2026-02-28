export interface ITimestampt {
  created_at: string;
  updated_at: string;
  deleted_at: null | string;
}

export interface IMember extends ITimestampt {
  id: number;
  full_name: string;
  phone_number: string;
  provider: string;
}

// base member-points
export interface IMemberPoint {
  id: number;
  member_id: number;
  total_points: number;
  updated_at: string | Date;
}

export interface IMemberWithMemberPoint extends IMember {
  member_point: IMemberPoint;
}
