package org.ndbs.spring.graalvm.product.domain.model;

import org.springframework.data.annotation.Id;
import org.springframework.data.relational.core.mapping.Table;

import java.util.Objects;

/**
 * Product class <br>
 *
 * <br>
 *
 * @author Nikolai Osipov <nao99.dev@gmail.com>
 * @since  2022-12-11
 */
@Table("products")
public class Product {
    @Id
    private final Long id;
    private final String category;
    private final String name;
    private final String description;

    private Product(Long id, String category, String name, String description) {
        this.id = id;
        this.category = category;
        this.name = name;
        this.description = description;
    }

    public static Product create(String category, String name, String description) {
        return new Product(null, category, name, description);
    }

    Product withId(Long id) {
        return new Product(id, category, name, description);
    }

    public Long getId() {
        return id;
    }

    public String getCategory() {
        return category;
    }

    public String getName() {
        return name;
    }

    public String getDescription() {
        return description;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) {
            return true;
        }

        if (o == null || getClass() != o.getClass()) {
            return false;
        }

        Product product = (Product) o;

        return Objects.equals(id, product.id)
            &&  Objects.equals(category, product.category)
            && Objects.equals(name, product.name);
    }

    @Override
    public int hashCode() {
        return Objects.hash(id, category, name);
    }
}
