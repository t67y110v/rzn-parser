CREATE  TABLE users (
  id serial not null primary key,
  email varchar not null unique,
  encrypted_password varchar not null,
  user_name varchar not null,
  seccond_name varchar not null,
  user_role varchar ,
  client_department boolean DEFAULT false , 
  education_department boolean DEFAULT false,
  source_tracking_department boolean DEFAULT false,
  periodic_reporting_department boolean DEFAULT false,
  international_department boolean DEFAULT false,
  documentation_department boolean DEFAULT false,
  nr_department boolean DEFAULT false,
  db_department boolean DEFAULT false, 
  monitoring_specialist boolean DEFAULT false,
  monitoring_responsible int 
);

