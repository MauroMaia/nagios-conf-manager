export interface Service {
  host_name: string;
  service_description: string;
  check_command: string;
  notifications_enabled: boolean;
  check_interval: number;
  retry_interval: number;
  use: string;
}
