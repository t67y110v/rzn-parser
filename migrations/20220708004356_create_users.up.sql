CREATE TABLE users (
  id serial not null primary key,
  email varchar not null unique,
  encrypted_password varchar not null,
  userName varchar not null,
  seccondName varchar not null,
  isadmin boolean DEFAULT false,
  educationdepartment boolean DEFAULT false,
  sourceTrackingdepartment boolean DEFAULT false,
  periodicReportingdepartment boolean DEFAULT false,
  internationaldepartment boolean DEFAULT false,
  documentationdepartment boolean DEFAULT false,
  nrdepartment boolean DEFAULT false,
  dbdepartment boolean DEFAULT false
);

