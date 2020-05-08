create table if not exists todo (
	id serial primary key,
	text varchar(200),
	done int
);
