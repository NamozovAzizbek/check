`creat table is not exists 'users'(
		'id' int(100) not null auto_increment,
		'username' varchar(45) not null,
		'password' varchar(40) not null,
		'isActive' tinyint(1) default null,
		primary key ('id'),
		unique key 'id_unique' ('id')
	)`