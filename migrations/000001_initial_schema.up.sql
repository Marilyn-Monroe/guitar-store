CREATE TABLE IF NOT EXISTS public.cart_item (
                                                user_id uuid NOT NULL,
                                                guitar_id uuid NOT NULL,
                                                quantity bigint NOT NULL
);

CREATE TABLE IF NOT EXISTS public.guitar (
                                             id uuid NOT NULL,
                                             name text NOT NULL,
                                             description text,
                                             sku text NOT NULL,
                                             price bigint NOT NULL,
                                             image text,
                                             type text NOT NULL,
                                             strings bigint NOT NULL,
                                             quantity_available bigint DEFAULT 0,
                                             created_at timestamp with time zone NOT NULL,
                                             modified_at timestamp with time zone NOT NULL,
                                             deleted_at timestamp with time zone
);

CREATE TABLE IF NOT EXISTS public.promocode (
                                                id uuid NOT NULL,
                                                name text NOT NULL,
                                                description text,
                                                code text NOT NULL,
                                                max_usage bigint,
                                                discount_amount bigint NOT NULL,
                                                expired_at timestamp with time zone,
                                                created_at timestamp with time zone NOT NULL,
                                                modified_at timestamp with time zone NOT NULL,
                                                deleted_at timestamp with time zone
);

CREATE TABLE IF NOT EXISTS public.review (
                                             id uuid NOT NULL,
                                             advantages text,
                                             disadvantages text,
                                             comments text,
                                             rating bigint NOT NULL,
                                             guitar_id uuid NOT NULL,
                                             created_at timestamp with time zone NOT NULL,
                                             created_by uuid NOT NULL
);

CREATE TABLE IF NOT EXISTS public."user" (
                                             id uuid NOT NULL,
                                             first_name text NOT NULL,
                                             last_name text NOT NULL,
                                             middle_name text NOT NULL,
                                             phone text NOT NULL,
                                             email text NOT NULL,
                                             created_at timestamp with time zone NOT NULL,
                                             last_login_at timestamp with time zone NOT NULL
);

ALTER TABLE ONLY public.guitar
    ADD CONSTRAINT guitar_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.promocode
    ADD CONSTRAINT promocode_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.review
    ADD CONSTRAINT review_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_pkey PRIMARY KEY (id);

INSERT INTO public.guitar (id, name, description, sku, price, image, type, strings, quantity_available, created_at, modified_at, deleted_at)
VALUES
    ('1e8b0c6b-dbd1-4d6f-b8be-e0c6fbbd6f2c', 'Fender Stratocaster', 'Classic electric guitar with a bright sound.', 'FEN12345', 120000, 'fender_strat.jpg', 'Electric', 6, 50, '2024-12-04T10:00:00+00:00', '2024-12-04T10:00:00+00:00', NULL),
    ('2a9c8b7a-f2f4-431f-8f27-c3f84deae06e', 'Gibson Les Paul', 'Solid body electric guitar known for its heavy tone.', 'GIB54321', 250000, 'gibson_lespaul.jpg', 'Electric', 6, 30, '2024-12-04T11:00:00+00:00', '2024-12-04T11:00:00+00:00', NULL),
    ('3b32e50f-c6b2-4a59-b060-3fe874b6583f', 'Yamaha Acoustic', 'Affordable acoustic guitar for beginners.', 'YAM67890', 30000, 'yamaha_acoustic.jpg', 'Acoustic', 6, 100, '2024-12-04T12:00:00+00:00', '2024-12-04T12:00:00+00:00', NULL),
    ('4f6780d2-5b5e-47fc-98b6-fc6b2e5d5b9e', 'Ibanez RG', 'High-performance electric guitar for rock musicians.', 'IBA11223', 150000, 'ibanez_rg.jpg', 'Electric', 7, 40, '2024-12-04T13:00:00+00:00', '2024-12-04T13:00:00+00:00', NULL),
    ('5b8e9e8e-7891-47ab-8bc2-60f74d18bfb0', 'Epiphone SG', 'Iconic electric guitar with a sharp sound.', 'EPI99876', 95000, 'epiphone_sg.jpg', 'Electric', 6, 25, '2024-12-04T14:00:00+00:00', '2024-12-04T14:00:00+00:00', NULL),
    ('6a5e7c12-f5c2-4ff1-b0a0-61adf415d0b3', 'Martin D-28', 'Premium acoustic guitar with a rich sound.', 'MAR12345', 500000, 'martin_d28.jpg', 'Acoustic', 6, 10, '2024-12-04T15:00:00+00:00', '2024-12-04T15:00:00+00:00', NULL),
    ('7c4a2e3a-5b21-4525-bd93-c7e497a1602b', 'PRS SE Custom', 'Electric guitar with a versatile tone.', 'PRS11234', 180000, 'prs_se_custom.jpg', 'Electric', 6, 60, '2024-12-04T16:00:00+00:00', '2024-12-04T16:00:00+00:00', NULL),
    ('8d3b6981-59d9-4166-b60b-6899b6a1e7e2', 'Taylor 214ce', 'Mid-range acoustic guitar with excellent sound quality.', 'TAY77654', 250000, 'taylor_214ce.jpg', 'Acoustic', 6, 35, '2024-12-04T17:00:00+00:00', '2024-12-04T17:00:00+00:00', NULL),
    ('9f9b5284-4369-4e72-9a01-d77b38a8f10a', 'Schecter C-1', 'Electric guitar designed for metal and hard rock.', 'SCH55432', 130000, 'schecter_c1.jpg', 'Electric', 7, 45, '2024-12-04T18:00:00+00:00', '2024-12-04T18:00:00+00:00', NULL),
    ('10e2a5c5-4e56-47d1-87f5-5d8204d6bc6f', 'Fender Jazz Bass', 'Electric bass guitar with a smooth, punchy sound.', 'FEN99876', 140000, 'fender_jazz_bass.jpg', 'Electric', 4, 20, '2024-12-04T19:00:00+00:00', '2024-12-04T19:00:00+00:00', NULL);

INSERT INTO public.review (id, advantages, disadvantages, comments, rating, guitar_id, created_at, created_by)
VALUES
    ('1b92a7d4-3bc1-4a7c-9b59-0cf7c6533b9d', 'Great tone, smooth playability, comfortable neck.', 'A bit heavy for long sessions.', 'I absolutely love the Fender Stratocaster. The bright sound is perfect for rock and blues.', 5, '1e8b0c6b-dbd1-4d6f-b8be-e0c6fbbd6f2c', '2024-12-04T10:15:00+00:00', '4f6f8f56-8cd9-4d3f-b3ad-033c6841d354'),
    ('2c22c9db-6df8-4b93-bbe5-0290a53b7cb6', 'Heavy tone, good sustain, classic design.', 'Quite expensive for the features it offers.', 'The Gibson Les Paul is a solid choice for those who enjoy a heavier tone. It feels like a premium guitar, though a bit pricier than some alternatives.', 4, '2a9c8b7a-f2f4-431f-8f27-c3f84deae06e', '2024-12-04T11:15:00+00:00', '5a3f8f1e-2ed9-47f9-918e-9d5310cd7485'),
    ('3d34c8a8-3509-4e3f-9272-88a49ad7d84c', 'Affordable, great sound for beginners, easy to play.', 'The finish feels a bit cheap.', 'I bought this Yamaha Acoustic as my first guitar, and it’s perfect for learning. The sound is rich for the price.', 4, '3b32e50f-c6b2-4a59-b060-3fe874b6583f', '2024-12-04T12:15:00+00:00', '1b2d3d4e-df98-45f4-9a02-e9a42e517f44'),
    ('4d21f7d7-2ea4-47d9-a4c4-74ff0f4d5769', 'Fast neck, great for shredding, solid build quality.', 'Not suitable for softer genres like jazz or acoustic.', 'The Ibanez RG is built for speed. It’s perfect for rock and metal, but might not suit other genres.', 4, '4f6780d2-5b5e-47fc-98b6-fc6b2e5d5b9e', '2024-12-04T13:15:00+00:00', '3f9a1b45-987e-44b7-8cc5-6d96f5a2c7d0'),
    ('5e13c7c3-6c9f-42da-8751-1a54b3744a37', 'Good resonance, vintage tone, affordable.', 'Not the best for very heavy music.', 'The Epiphone SG gives me a vintage vibe with its classic sound. Great for a variety of rock styles.', 4, '5b8e9e8e-7891-47ab-8bc2-60f74d18bfb0', '2024-12-04T14:15:00+00:00', '7d31b5b4-2e9a-4c7e-8467-b4bc1a27388c'),
    ('6f3a6b6b-9637-4088-9d2b-09a911d0f973', 'Rich tone, beautiful craftsmanship, versatile.', 'Expensive, not ideal for beginners.', 'The Martin D-28 is a work of art. The tone is amazing, but the price is a little steep for someone just starting.', 5, '6a5e7c12-f5c2-4ff1-b0a0-61adf415d0b3', '2024-12-04T15:15:00+00:00', '4e0e64cf-66d6-4c5a-b99d-5e65172fa3fc'),
    ('7a84b30d-83f9-4a35-9089-8b6cc9f4591a', 'Great for all genres, comfortable neck, durable.', 'Price could be lower for the value.', 'PRS SE Custom has excellent build quality, and its versatile tone is perfect for many musical styles. A great guitar overall!', 4, '7c4a2e3a-5b21-4525-bd93-c7e497a1602b', '2024-12-04T16:15:00+00:00', '2e8f3f47-dfa1-4f82-b3d7-1a2749a72a92'),
    ('8b65e23d-3c9b-4ed7-b499-e95b09f4e522', 'Warm tone, easy to play, beautiful design.', 'The sound could be louder in a live performance setting.', 'Taylor 214ce is a fantastic choice for anyone who loves acoustic guitars. It feels natural to play, and the tone is very warm.', 4, '8d3b6981-59d9-4166-b60b-6899b6a1e7e2', '2024-12-04T17:15:00+00:00', '54ac9c74-b3b6-4e1f-a1e4-48d6a0d3fe0f'),
    ('9c8e9201-76c4-4b5f-9aeb-2de7b8723b88', 'Affordable, solid build, sharp high-end sound.', 'Lacks the depth of more expensive models.', 'The Schecter C-1 is an excellent option for players who want an affordable guitar with solid features. It’s perfect for rock and metal.', 4, '9f9b5284-4369-4e72-9a01-d77b38a8f10a', '2024-12-04T18:15:00+00:00', '1a63c05c-e8db-4c59-b7e9-bc8bc4e5c75e'),
    ('10d9c7f3-55d3-4f6e-91b0-4fdb5b9a6c1e', 'Great for beginners, comfortable and easy to use.', 'Doesn’t have the same range as higher-end models.', 'The Fender Jazz Bass is a great choice for those just starting out with bass guitar. It’s easy to play and sounds great.', 4, '10e2a5c5-4e56-47d1-87f5-5d8204d6bc6f', '2024-12-04T19:15:00+00:00', '5b37c99d-2041-4b1f-9226-5f4edc7599f6');

INSERT INTO public."user" (id, first_name, last_name, middle_name, phone, email, created_at, last_login_at)
VALUES
    ('f9b8a5e1-214b-4d62-aeab-c5e02f2fa60b', 'John', 'Doe', 'Michael', '+1234567890', 'john.doe@example.com', '2024-01-15T08:00:00+00:00', '2024-12-04T09:00:00+00:00'),
    ('fbb04c3a-3c55-4200-a9a9-0fcf92e0566b', 'Jane', 'Smith', 'Marie', '+1234567891', 'jane.smith@example.com', '2024-02-10T09:30:00+00:00', '2024-12-04T09:30:00+00:00'),
    ('b9d2f5b1-76d1-4a60-b9ec-23823cd4d42f', 'Alice', 'Johnson', 'Rose', '+1234567892', 'alice.johnson@example.com', '2024-03-05T10:00:00+00:00', '2024-12-04T10:00:00+00:00'),
    ('8e4f2e98-bb92-4f76-8a7b-4f0f515d58a1', 'Bob', 'Brown', 'William', '+1234567893', 'bob.brown@example.com', '2024-04-20T11:00:00+00:00', '2024-12-04T10:30:00+00:00'),
    ('9a6b0371-3f64-4c80-a56f-6e7a682396f2', 'Eve', 'White', 'Lily', '+1234567894', 'eve.white@example.com', '2024-05-25T12:00:00+00:00', '2024-12-04T11:00:00+00:00'),
    ('d15c7065-8a0d-4520-a0ac-8acac1e0fcf3', 'Chris', 'Davis', 'Lee', '+1234567895', 'chris.davis@example.com', '2024-06-30T13:00:00+00:00', '2024-12-04T11:30:00+00:00'),
    ('b12e48f1-32f9-4d9f-9a14-17ed5e432823', 'Daniel', 'Miller', 'George', '+1234567896', 'daniel.miller@example.com', '2024-07-10T14:00:00+00:00', '2024-12-04T12:00:00+00:00'),
    ('f16e7c28-d3d0-47e0-94a9-b0b5d63784c9', 'Olivia', 'Wilson', 'Grace', '+1234567897', 'olivia.wilson@example.com', '2024-08-15T15:00:00+00:00', '2024-12-04T12:30:00+00:00'),
    ('c55c50a7-f7c9-4e79-9b56-d53160f7d2a4', 'James', 'Taylor', 'Henry', '+1234567898', 'james.taylor@example.com', '2024-09-20T16:00:00+00:00', '2024-12-04T13:00:00+00:00'),
    ('3fe923ba-d0ae-4f2d-9d7c-88e6f041987d', 'Sophia', 'Martinez', 'Sophia', '+1234567899', 'sophia.martinez@example.com', '2024-10-25T17:00:00+00:00', '2024-12-04T13:30:00+00:00');

INSERT INTO public.promocode (id, name, description, code, max_usage, discount_amount, expired_at, created_at, modified_at, deleted_at)
VALUES
    ('3e6a9a74-8f67-4cd7-9448-c118ac935e39', 'Winter Sale', 'Get 20% off on all guitars during the winter sale!', 'WINTER20', 500, 2000, '2024-12-31T23:59:59+00:00', '2024-12-01T09:00:00+00:00', '2024-12-01T09:00:00+00:00', NULL),
    ('e1ffdb54-b602-4b97-a6d3-64334c0fd46a', 'New Year Special', 'Enjoy a 15% discount on your first purchase in the new year!', 'NEWYEAR15', 300, 1500, '2025-01-15T23:59:59+00:00', '2024-12-05T10:00:00+00:00', '2024-12-05T10:00:00+00:00', NULL),
    ('c5b2159c-2539-4ff5-9887-f4e8fe6da039', 'Black Friday Deal', '50% off on selected electric guitars for Black Friday!', 'BLACKFRI50', 100, 5000, '2024-11-29T23:59:59+00:00', '2024-11-10T08:30:00+00:00', '2024-11-10T08:30:00+00:00', NULL),
    ('a00b98a8-79b2-4b28-b45c-2b1d2bc5b9b9', 'Summer Discount', 'Get a 10% discount on all orders over $100!', 'SUMMER10', 1000, 1000, '2024-08-31T23:59:59+00:00', '2024-06-01T09:00:00+00:00', '2024-06-01T09:00:00+00:00', NULL),
    ('f7b6cc71-fac6-44ab-88db-0f7acb0f518b', 'Holiday Special', 'Exclusive 25% off on premium guitars for the holidays!', 'HOLIDAY25', 200, 2500, '2024-12-25T23:59:59+00:00', '2024-11-25T12:00:00+00:00', '2024-11-25T12:00:00+00:00', NULL);
