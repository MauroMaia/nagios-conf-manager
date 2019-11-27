export interface Contact
{

  name: string;

  contact_name: string;

  use: string;

  alias: string;

  email: string;

  host_notifications_enabled: boolean;

  service_notifications_enabled: boolean;

  host_notification_period: string;

  service_notification_period: string;

  host_notification_commands: string;

  service_notification_commands: string;

  host_notification_options: string[];

  service_notification_options: string[];

  IsTemplate: boolean;
}
