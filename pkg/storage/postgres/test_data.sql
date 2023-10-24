TRUNCATE TABLE posts;

INSERT INTO posts (id, title, content, post_time, link) VALUES
(1, 'Title 1', 'Content 1', EXTRACT(EPOCH FROM NOW()), 'https://gonews.xyz'),
(2, 'Title 2', 'Content 2', EXTRACT(EPOCH FROM NOW()), 'https://gonews.xyz'),
(3, 'Title 3', 'Content 3', EXTRACT(EPOCH FROM NOW()), 'https://gonews.xyz'),
(4, 'Title 4', 'Content 4', EXTRACT(EPOCH FROM NOW()), 'https://gonews.xyz'),
(5, 'Title 5', 'Content 5', EXTRACT(EPOCH FROM NOW()), 'https://gonews.xyz'),
(6, 'Title 6', 'Content 6', EXTRACT(EPOCH FROM NOW()), 'https://gonews.xyz'),
(7, 'Title 7', 'Content 7', EXTRACT(EPOCH FROM NOW()), 'https://gonews.xyz'),
(8, 'Title 8', 'Content 8', EXTRACT(EPOCH FROM NOW()), 'https://gonews.xyz'),
(9, 'Title 9', 'Content 9', EXTRACT(EPOCH FROM NOW()), 'https://gonews.xyz'),
(10, 'Title 10', 'Content 10', EXTRACT(EPOCH FROM NOW()), 'https://gonews.xyz'),
(11, 'Title 11', 'Content 11', EXTRACT(EPOCH FROM NOW()), 'https://gonews.xyz'),
(12, 'Title 12', 'Content 12', EXTRACT(EPOCH FROM NOW()), 'https://gonews.xyz'),
(13, 'Title 13', 'Content 13', EXTRACT(EPOCH FROM NOW()), 'https://gonews.xyz'),
(14, 'Title 14', 'Content 14', EXTRACT(EPOCH FROM NOW()), 'https://gonews.xyz'),
(15, 'Title 15', 'Content 15', EXTRACT(EPOCH FROM NOW()), 'https://gonews.xyz');