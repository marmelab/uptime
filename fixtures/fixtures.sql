INSERT INTO destination (destination)
VALUES 
('website'),
('google.fr'),
('youtube.fr'),
('yahoo.fr'),
('lolking.com'),
('thisWebsiteDoesntExist.fr');
INSERT INTO results (target_id, destination, status, duration)
VALUES
(1, 'google.fr', 'good',2115),
(2, 'website', 'failed',-1),
(3, 'youtube.fr', 'good', 300),
(4, 'yahoo.fr', 'good', 300),
(5, 'lolking.com', 'good', 300),
(6, 'thisWebsiteDoesntExist.fr', 'false', 300);
