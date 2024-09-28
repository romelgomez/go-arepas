-- CreateTable
CREATE TABLE "Post" (
    "id" VARCHAR(100) NOT NULL DEFAULT gen_random_uuid(),
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3),
    "title" VARCHAR(100) NOT NULL,
    "sub_title" VARCHAR(100),
    "published" BOOLEAN NOT NULL DEFAULT false,
    "description" VARCHAR(100),

    CONSTRAINT "Post_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "Post_id_key" ON "Post"("id");
