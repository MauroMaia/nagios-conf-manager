export interface Host {
  host_name: string;
  name: string;
  alias: string;
  address: string;
  max_check_attempts: number;
  notification_interval: number;
  notification_period: string;
  contact: string;
  contact_groups: string;
  check_period: string;
  register: boolean;
  use: string;
}
