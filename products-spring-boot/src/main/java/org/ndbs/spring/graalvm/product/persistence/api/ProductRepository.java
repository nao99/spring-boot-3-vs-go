package org.ndbs.spring.graalvm.product.persistence.api;

import org.ndbs.spring.graalvm.product.domain.model.Product;
import org.springframework.data.jdbc.repository.query.Query;
import org.springframework.data.repository.CrudRepository;

import java.util.Set;

/**
 * ProductRepository interface <br>
 *
 * <br>
 *
 * @author Nikolai Osipov <nao99.dev@gmail.com>
 * @since  2022-12-11
 */
public interface ProductRepository extends CrudRepository<Product, Long> {
    @Query("SELECT * FROM products ORDER BY id DESC LIMIT 20;")
    Set<Product> findLast20();
}
