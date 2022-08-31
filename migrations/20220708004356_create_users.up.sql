CREATE TABLE users (
  id serial not null primary key,
  email varchar not null unique,
  encrypted_password varchar not null,\
  userName varchar not null,
  seccondName varchar not null,
  isadmin boolean DEFAULT false,
  educationDepartment boolean DEFAULT false,
  sourceTrackingDepartment boolean DEFAULT false,
  periodicReportingDepartment boolean DEFAULT false,
  internationalDepartment boolean DEFAULT false,
  documentationDepartment boolean DEFAULT false,
  nrDepartment boolean DEFAULT false,
  dbDepartment boolean DEFAULT false
);